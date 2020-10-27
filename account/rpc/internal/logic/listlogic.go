package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &account.AccountList{}, nil
}
