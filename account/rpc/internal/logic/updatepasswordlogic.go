package logic

import (
	"account/crypto"
	"context"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"account/rpc/internal/svc"
	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(in *account.UpdatePasswordRequest) (*account.AccountEmpty, error) {
	if in.Uuid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid uuid")
	}
	if len(in.Password) < l.svcCtx.MinPasswordLength {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid password - it must be at least %d characters long", l.svcCtx.MinPasswordLength)
	}
	salt, err := crypto.NewSalt()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to generate a salt: %s", err)
	}
	hash, err := crypto.HashPassword(salt, []byte(in.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to hash the password: %s", err)
	}
	ret, err := l.svcCtx.Model.UpdatePassword(in.Uuid, string(hash), string(salt))
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to update the password: %s", err)
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to read the database: %s", err)
	}
	if n != 1 {
		return nil, status.Errorf(codes.NotFound, "")
	}
	return &account.AccountEmpty{}, nil
}
