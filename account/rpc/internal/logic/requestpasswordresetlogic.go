package logic

import (
	"context"

	"account/rpc/internal/svc"
	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type RequestPasswordResetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRequestPasswordResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestPasswordResetLogic {
	return &RequestPasswordResetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RequestPasswordResetLogic) RequestPasswordReset(in *account.PasswordResetRequest) (*account.AccountEmpty, error) {
	// todo: add your logic here and delete this line

	return &account.AccountEmpty{}, nil
}
