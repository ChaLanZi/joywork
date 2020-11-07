package logic

import (
	"account/crypto"
	"account/rpc/internal/auth"
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

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *account.CreateAccountRequest) (*account.Account, error) {
	_, authz, err := helper.GetAuth(l.ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to authorize.")
	}
	switch authz {
	case auth.AuthorizationSupportUser:
	case auth.AuthorizationWWWService:
	case auth.AuthorizationCompanyService:
	default:
		return nil, status.Errorf(codes.PermissionDenied, "You do not have access to this service.")
	}
	if (len(in.Email) + len(in.PhoneNumber) + len(in.Name)) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Empty request")
	}
	if in.PhoneNumber, err = helper.ParseAndFormatPhoneNumber(in.PhoneNumber); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid phone number")
	}
	in.Email = strings.TrimSpace(strings.ToLower(in.Email))
	if len(in.Email) > 0 && strings.Index(in.Email, "@") == -1 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid email")
	}

	var a *model.Account
	if in.Email != "" {
		a, err = l.svcCtx.Model.FindAccountByEmail(in.Email)
		if err == nil {
			return &account.Account{
				Uuid:               a.Id,
				Name:               a.Name,
				Email:              a.Email,
				PhoneNumber:        a.PhoneNumber,
				ProtoUrl:           a.PhotoUrl,
				ConfirmedAndActive: helper.Int64ToBool(a.ConfirmedAndActive),
				Support:            helper.Int64ToBool(a.Support),
			}, status.Errorf(codes.AlreadyExists, "A user with that email already exists,Try a password reset.")
		}
		if err != model.ErrNotFound {
			return nil, status.Errorf(codes.Unknown, "An unknown error occurred while searching for that email: %s", err)
		}
	}
	if in.PhoneNumber != "" {
		a, err = l.svcCtx.Model.FindAccountByPhoneNumber(in.PhoneNumber)
		if err == nil {
			return &account.Account{
				Uuid:               a.Id,
				Name:               a.Name,
				Email:              a.Email,
				PhoneNumber:        a.PhoneNumber,
				ProtoUrl:           a.PhotoUrl,
				ConfirmedAndActive: helper.Int64ToBool(a.ConfirmedAndActive),
				Support:            helper.Int64ToBool(a.Support),
			}, status.Errorf(codes.AlreadyExists, "A user with that phone number already exists,Try a password reset.")
		}

		if err != model.ErrNotFound {
			return nil, status.Errorf(codes.Unknown, "An unknown error occurred while searching for that phone number: %s", err)
		}
	}
	uuid, err := crypto.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "%s", err)
	}
	a.Id = uuid.String()
	a.Name = in.Name
	a.PhoneNumber = in.PhoneNumber
	a.Email = in.Email
	a.PhotoUrl = helper.GenerateGravatarURL(in.Email)
	a.ConfirmedAndActive = 0
	a.Support = 0

	_, err = l.svcCtx.Model.Insert(a)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed Insert database: %s", err)
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
