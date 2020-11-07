package svc

import (
	"account/rpc/internal/config"
	"account/rpc/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c                 config.Config
	Model             *model.AccountModel
	MinPasswordLength int
	SigningSecret     string
	ExternalApex      string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:                 c,
		MinPasswordLength: c.MinPasswordLength,
		SigningSecret:     c.SigningSecret,
		ExternalApex:      c.ExternalApex,
		Model:             model.NewAccountModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table),
	}
}
