package logic

import (
	"context"

	"account/rpc/internal/svc"
	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangeEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeEmailLogic {
	return &ChangeEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeEmailLogic) ChangeEmail(in *account.EmailComfirmation) (*account.AccountEmpty, error) {
	// todo: add your logic here and delete this line

	return &account.AccountEmpty{}, nil
}
