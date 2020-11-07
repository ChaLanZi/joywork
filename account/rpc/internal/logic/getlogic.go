package logic

import (
	"account/rpc/internal/auth"
	"account/rpc/internal/helper"
	"account/rpc/model"
	"context"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"account/rpc/internal/svc"
	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *account.GetAccountRequest) (*account.Account, error) {
	md, authz, err := helper.GetAuth(l.ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to authorize.")
	}
	switch authz {
	case auth.AuthorizationWWWService:
	case auth.AuthorizationAccountService:
	case auth.AuthorizationCompanyService:
	case auth.AuthorizationWhoamiService:
	case auth.AuthorizationBotService:
	case auth.AuthorizationAuthenticatedUser:
		uuid, err := auth.GetCurrentUserUUIDFromMetadata(md)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "Failed to find current user uuid %v", md)
		}
		if uuid != in.Uuid {
			return nil, status.Errorf(codes.PermissionDenied, "You do not have access to this service.")
		}
	case auth.AuthorizationSupportUser:
	case auth.AuthorizationSuperpowersService:
	default:
		return nil, status.Errorf(codes.PermissionDenied, "You do not have access to this service.")
	}
	if in.Uuid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid uuid")
	}
	a, err := l.svcCtx.Model.FindOne(in.Uuid)
	if err == model.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "A user not exist")
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to find database: ", err)
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
