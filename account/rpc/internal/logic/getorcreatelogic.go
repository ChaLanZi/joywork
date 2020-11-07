package logic

import (
	"account/crypto"
	"account/rpc/internal/helper"
	"account/rpc/model"
	"context"
	"strings"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

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
	var err error
	in.Email = strings.ToLower(in.Email)
	if in.PhoneNumber, err = helper.ParseAndFormatPhoneNumber(in.PhoneNumber); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, " Invalid phone number")
	}
	var a *model.Account
	if len(in.Email) > 0 {
		a, err = l.svcCtx.Model.FindAccountByEmail(in.Email)
		if err != nil && err != model.ErrNotFound {
			return nil, status.Errorf(codes.Unknown, "failed to query database for existing email: ", err)
		}
	}
	if len(in.PhoneNumber) > 0 && a == nil {
		a, err = l.svcCtx.Model.FindAccountByPhoneNumber(in.PhoneNumber)
		if err != nil && err != model.ErrNotFound {
			return nil, status.Errorf(codes.Unknown, "failed to query database for existing phone number: ", err)
		}
	}

	if a == nil {
		uuid, err := crypto.NewUUID()
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "", err)
		}
		a.Id = uuid.String()
		a.Name = in.Name
		a.Email = in.Email
		a.PhoneNumber = in.PhoneNumber
		a.PhotoUrl = helper.GenerateGravatarURL(in.Email)
		a.ConfirmedAndActive = 0
		a.Support = 0
		_, err = l.svcCtx.Model.Insert(a)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, " failed to insert database: ", err)
		}
	}

	return &account.Account{
		Uuid:               a.Id,
		Name:               a.Name,
		Email:              a.Email,
		PhoneNumber:        a.PhoneNumber,
		ProtoUrl:           a.PhotoUrl,
		ConfirmedAndActive: helper.Int64ToBool(a.ConfirmedAndActive),
		Support:            helper.Int64ToBool(a.Support),
	}, nil
}
