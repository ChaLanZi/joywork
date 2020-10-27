package logic

import (
	"context"

	"account/rpc/internal/svc"
	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type TrackEventLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTrackEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrackEventLogic {
	return &TrackEventLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TrackEventLogic) TrackEvent(in *account.TrackEventReqeust) (*account.AccountEmpty, error) {
	// todo: add your logic here and delete this line

	return &account.AccountEmpty{}, nil
}
