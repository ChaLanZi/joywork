package logic

import (
	"account/rpc/helper"
	"account/rpc/model"
	"context"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"account/rpc/internal/svc"
	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAccountByPhoneNumberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountByPhoneNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountByPhoneNumberLogic {
	return &GetAccountByPhoneNumberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountByPhoneNumberLogic) GetAccountByPhoneNumber(in *account.GetAccountByPhoneNumberRequest) (*account.Account, error) {
	var err error
	if in.PhoneNumber, err = helper.ParseAndFormatPhoneNumber(in.PhoneNumber); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid phone number")
	}
	if in.PhoneNumber == "" {
		return nil, status.Errorf(codes.InvalidArgument, "No Phone number provided")
	}
	var a *model.Account
	a, err = l.svcCtx.Model.FindAccountByPhoneNumber(in.PhoneNumber)
	if err == model.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "")
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to query database for existing phone number")
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
