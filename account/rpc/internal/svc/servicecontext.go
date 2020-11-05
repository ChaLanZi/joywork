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
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:                 c,
		MinPasswordLength: c.MinPasswordLength,
		Model:             model.NewAccountModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table),
	}
}
