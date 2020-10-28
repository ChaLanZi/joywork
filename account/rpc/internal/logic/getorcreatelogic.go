package logic

import (
	"context"

	"account/rpc/internal/svc"

	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOrCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrCreateLogic {
	return &GetOrCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrCreateLogic) GetOrCreate(in *account.GetOrCreateRequest) (*account.Account, error) {
	// todo: add your logic here and delete this line

	return &account.Account{}, nil
}
