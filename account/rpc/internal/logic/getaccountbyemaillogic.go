package logic

import (
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

type GetAccountByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountByEmailLogic {
	return &GetAccountByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountByEmailLogic) GetAccountByEmail(in *account.GetAccountByEmailRequest) (*account.Account, error) {
	var err error
	in.Email = strings.ToLower(in.Email)
	if in.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid email")
	}
	var a *model.Account
	a, err = l.svcCtx.Model.FindAccountByEmail(in.Email)
	if err == model.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "")
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to query database for existing email")
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
