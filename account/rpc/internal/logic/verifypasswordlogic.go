package logic

import (
	"account/crypto"
	"account/rpc/internal/auth"
	"account/rpc/internal/helper"
	"account/rpc/model"
	"context"
	"strings"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"account/rpc/internal/svc"
	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerifyPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyPasswordLogic {
	return &VerifyPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyPasswordLogic) VerifyPassword(in *account.VerifyPasswordRequest) (*account.Account, error) {
	_, authz, err := helper.GetAuth(l.ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to authorize.")
	}
	switch authz {
	case auth.AuthorizationWWWService:
	case auth.AuthorizationSupportUser:
	default:
		return nil, status.Errorf(codes.PermissionDenied, "You do not have access to this service.")
	}
	in.Email = strings.TrimSpace(strings.ToLower(in.Email))
	in.Password = strings.TrimSpace(in.Password)
	if len(in.Email) <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid email ")
	}
	if len(in.Email) > 0 && strings.Index(in.Email, "@") == -1 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid email ")
	}
	if len(in.Password) < l.svcCtx.MinPasswordLength {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid password")
	}
	a, err := l.svcCtx.Model.FindAccountByEmail(in.Email)
	if err == model.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "")
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to find the database by email: %S", err)
	}
	if a.ConfirmedAndActive == 0 {
		return nil, status.Errorf(codes.PermissionDenied, "This user has not confirmed their account")
	}
	if len(a.PasswordHash) == 0 {
		return nil, status.Errorf(codes.PermissionDenied, "This user has not set up their password ")
	}
	if crypto.CheckPasswordHash([]byte(a.PasswordHash), []byte(a.PasswordSalt), []byte(in.Password)) != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Incorrect password")
	}

	return &account.Account{
		Uuid:               a.Id,
		Name:               a.Name,
		Email:              a.Email,
		PhoneNumber:        a.PhoneNumber,
		ProtoUrl:           a.PhotoUrl,
		ConfirmedAndActive: helper.Int64ToBool(a.ConfirmedAndActive),
		Support:            helper.Int64ToBool(a.Support),
	}, nil
}
