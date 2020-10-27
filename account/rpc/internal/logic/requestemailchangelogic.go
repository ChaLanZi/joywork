package logic

import (
	"context"

	account "account/rpc/pb"

	"account/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type RequestEmailChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRequestEmailChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestEmailChangeLogic {
	return &RequestEmailChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RequestEmailChangeLogic) RequestEmailChange(in *account.EmailChangeRequest) (*account.AccountEmpty, error) {
	// todo: add your logic here and delete this line

	return &account.AccountEmpty{}, nil
}
