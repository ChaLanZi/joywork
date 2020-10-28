package logic

import (
	"context"

	account "account/rpc/pb"

	"account/rpc/internal/svc"

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

func (l *ChangeEmailLogic) ChangeEmail(in *account.EmailChangeRequest) (*account.AccountEmpty, error) {
	// todo: add your logic here and delete this line

	return &account.AccountEmpty{}, nil
}