package logic

import (
	"account/rpc/model"
	"context"
	"strings"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

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
	var err error
	in.Email = strings.TrimSpace(strings.ToLower(in.Email))
	if in.Uuid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid uuid")
	}
	if in.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid email")
	}

	_, err = l.svcCtx.Model.FindAccountByEmail(in.Email)
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "A user that email already exists.")
	}
	if err != model.ErrNotFound {
		return nil, status.Errorf(codes.Unknown, "failed to find database by email: %s", err)
	}

	res, err := l.svcCtx.Model.UpdateEmail(in.Uuid, in.Email)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to update the database: %s ", err)
	}
	row, err := res.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to read the database: %s ", err)
	}
	if row != 1 {
		return nil, status.Errorf(codes.NotFound, "")
	}

	return &account.AccountEmpty{}, nil
}
