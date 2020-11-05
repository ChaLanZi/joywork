package logic

import (
	"context"

	"account/rpc/internal/svc"
	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type SyncUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncUserLogic {
	return &SyncUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SyncUserLogic) SyncUser(in *account.SyncUserRequest) (*account.AccountEmpty, error) {

	return &account.AccountEmpty{}, nil
}
