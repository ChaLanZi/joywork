package logic

import (
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

type RequestPasswordResetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRequestPasswordResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestPasswordResetLogic {
	return &RequestPasswordResetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RequestPasswordResetLogic) RequestPasswordReset(in *account.PasswordResetRequest) (*account.AccountEmpty, error) {
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
	if len(in.Email) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "No Email provided")
	}

	a, err := l.svcCtx.Model.FindAccountByEmail(in.Email)
	if err == model.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "No user with that email exists.")
	} else if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to find database by email: %v", err)
	}
	if a == nil {
		return nil, status.Errorf(codes.NotFound, "No user with that email exists.")
	}
	return &account.AccountEmpty{}, nil
}
