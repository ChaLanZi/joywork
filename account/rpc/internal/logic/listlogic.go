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

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *account.GetAccountListRequest) (*account.AccountList, error) {
	_, authz, err := helper.GetAuth(l.ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to authorize.")
	}
	switch authz {
	case auth.AuthorizationSupportUser:
	default:
		return nil, status.Errorf(codes.PermissionDenied, "You do not have access to this service.")
	}
	if in.Offset < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid offset - must be greater than or equal to zero")
	}
	if in.Limit <= 0 {
		in.Limit = 10
	}
	rows, err := l.svcCtx.Model.FindAll(in.Offset, in.Limit)
	if err == model.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "User not exist")
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "")
	}
	var accounts account.AccountList
	for _, r := range rows {
		a := &account.Account{
			Uuid:               r.Id,
			Name:               r.Name,
			Email:              r.Email,
			PhoneNumber:        r.PhoneNumber,
			ProtoUrl:           r.PhotoUrl,
			ConfirmedAndActive: helper.Int64ToBool(r.ConfirmedAndActive),
			Support:            helper.Int64ToBool(r.Support),
		}
		accounts.Accounts = append(accounts.Accounts, a)

	}
	return &accounts, nil
}
