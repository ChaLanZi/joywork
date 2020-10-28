package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &account.Account{}, nil
}
