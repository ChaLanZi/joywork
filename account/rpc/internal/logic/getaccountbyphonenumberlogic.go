package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &account.Account{}, nil
}
