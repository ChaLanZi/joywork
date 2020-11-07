package logic

import (
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

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *account.Account) (*account.Account, error) {
	md, authz, err := helper.GetAuth(l.ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to authorize")
	}
	switch authz {
	case auth.AuthorizationWWWService:
	case auth.AuthorizationCompanyService:
	case auth.AuthorizationAuthenticatedUser:
		uuid, err := auth.GetCurrentUserUUIDFromMetadata(md)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "failed to find current user uuid")
		}
		if uuid != in.Uuid {
			return nil, status.Errorf(codes.PermissionDenied, "You do not have access to this service.")
		}
	case auth.AuthorizationSupportUser:
	case auth.AuthorizationSuperpowersService:
	default:
		return nil, status.Errorf(codes.PermissionDenied, "You do not have access to this service.")
	}
	if in.Uuid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid uuid")
	}
	a, err := l.svcCtx.Model.FindOne(in.Uuid)
	if err == model.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "A user not exist.")
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to find database: ", err)
	}
	if in.MemberSince != a.MemberSince.Unix() {
		return nil, status.Errorf(codes.PermissionDenied, "You cannot modify the member_since date.")
	}
	in.Email = strings.TrimSpace(strings.ToLower(in.Email))
	if in.Email != "" && in.Email != a.Email {
		_, err := l.svcCtx.Model.FindAccountByEmail(in.Email)
		if err == nil {
			return nil, status.Errorf(codes.AlreadyExists, "A user that email already exists.")
		}
		if err != model.ErrNotFound {
			return nil, status.Errorf(codes.Unknown, "failed to find database by email:", err)
		}
	}
	if in.PhoneNumber, err = helper.ParseAndFormatPhoneNumber(in.PhoneNumber); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid phone number")
	}
	if in.PhoneNumber != "" && in.PhoneNumber != a.PhoneNumber {
		_, err := l.svcCtx.Model.FindAccountByPhoneNumber(in.PhoneNumber)
		if err == nil {
			return nil, status.Errorf(codes.AlreadyExists, "A user that phone umber already exists.")
		}

		if err != model.ErrNotFound {
			return nil, status.Errorf(codes.Unknown, "failed to find database by phone number: ", err)
		}
	}
	if authz == auth.AuthorizationAuthenticatedUser {
		if in.ConfirmedAndActive != helper.Int64ToBool(a.ConfirmedAndActive) && a.ConfirmedAndActive == 0 {
			return nil, status.Errorf(codes.PermissionDenied, "You cannot activate this account.")
		}
		if in.Support != helper.Int64ToBool(a.Support) {
			return nil, status.Errorf(codes.PermissionDenied, "You cannot change the support parameter.")
		}
		if in.ProtoUrl != in.ProtoUrl {
			return nil, status.Errorf(codes.PermissionDenied, "You cannot change the photo through this endpoint.")
		}
		if in.Email != a.Email {
			emailChange := NewRequestEmailChangeLogic(l.ctx, l.svcCtx)
			_, err = emailChange.RequestEmailChange(&account.EmailChangeRequest{Uuid: in.Uuid, Email: in.Email})
			if err != nil {
				return nil, status.Errorf(codes.Unknown, "failed to request email change.")
			}
			// revert
			in.Email = a.Email
		}
	}
	res, err := l.svcCtx.Model.Update(model.Account{
		Name:               in.Name,
		Email:              in.Email,
		PhoneNumber:        in.PhoneNumber,
		PhotoUrl:           helper.GenerateGravatarURL(in.Email),
		ConfirmedAndActive: helper.BoolToInt64(in.ConfirmedAndActive),
		Support:            helper.BoolToInt64(in.Support),
	})
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to update the database:", err)
	}
	r, err := res.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to read the database:", err)
	}
	if r != 1 {
		return nil, status.Errorf(codes.NotFound, "")
	}
	return &account.Account{}, nil
}
