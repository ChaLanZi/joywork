package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	accountFieldNames        = builderx.FieldNames(&Account{})
	accountRows              = strings.Join(accountFieldNames, ",")
	accountRowsExpectAutoSet = strings.Join(stringx.Remove(accountFieldNames, "create_time", "update_time"), ",")
	// accountRowsWithPlaceHolder                     = strings.Join(stringx.Remove(accountFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
	accountRowsWithPlaceAndRemoveMemberSinceHolder = strings.Join(stringx.Remove(accountFieldNames, "id", "member_since", "create_time", "update_time"), "=?,") + "=?"

	cacheAccountIdPrefix = "cache#Account#id#"
)

type (
	AccountModel struct {
		sqlc.CachedConn
		table string
	}

	Account struct {
		Id                 string    `db:"id"`
		Name               string    `db:"name"`
		Email              string    `db:"email"`
		PhoneNumber        string    `db:"phone_number"`
		ConfirmedAndActive int64     `db:"confirmed_and_active"`
		MemberSince        time.Time `db:"member_since"`
		PasswordHash       string    `db:"password_hash"`
		PasswordSalt       string    `db:"password_salt"`
		PhotoUrl           string    `db:"photo_url"`
		Support            int64     `db:"support"`
	}
)

func NewAccountModel(conn sqlx.SqlConn, c cache.CacheConf, table string) *AccountModel {
	return &AccountModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      table,
	}
}

func (m *AccountModel) Insert(data *Account) (sql.Result, error) {
	query := `insert into ` + m.table + ` (` + accountRowsExpectAutoSet + `) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	ret, err := m.ExecNoCache(query, data.Id, data.Name, data.Email, data.PhoneNumber, data.ConfirmedAndActive, data.MemberSince, data.PasswordHash, data.PasswordSalt, data.PhotoUrl, data.Support)

	return ret, err
}

func (m *AccountModel) FindOne(id string) (*Account, error) {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, id)
	var resp Account
	err := m.QueryRow(&resp, accountIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select ` + accountRows + ` from ` + m.table + ` where id = ? limit 1`
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *AccountModel) FindAll(offset, limit int32) ([]*Account, error) {
	var rows []*Account
	query := `select ` + accountRows + ` from ` + m.table + ` limit=? offset=? `
	err := m.QueryRowsNoCache(&rows, query, limit, offset)
	switch err {
	case nil:
		return rows, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *AccountModel) FindAccountByEmail(email string) (*Account, error) {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, email)
	var resp Account
	err := m.QueryRow(&resp, accountIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select ` + accountRows + ` from ` + m.table + ` where email=? limit=1`
		return conn.QueryRow(v, query, email)
	})

	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *AccountModel) FindAccountByPhoneNumber(phoneNumber string) (*Account, error) {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, phoneNumber)
	var resp Account
	err := m.QueryRow(&resp, accountIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select ` + accountRows + ` from ` + m.table + ` where phone_Number = ? limit=1 `
		return conn.QueryRow(v, query, phoneNumber)
	})

	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *AccountModel) Update(data Account) (sql.Result, error) {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, data.Id)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `update ` + m.table + ` set ` + accountRowsWithPlaceAndRemoveMemberSinceHolder + ` where id = ?`
		return conn.Exec(query, data.Name, data.Email, data.PhoneNumber, data.ConfirmedAndActive, data.PasswordHash, data.PasswordSalt, data.PhotoUrl, data.Support, data.Id)
	}, accountIdKey)
	return ret, err
}

func (m *AccountModel) UpdateEmail(uuid, email string) (sql.Result, error) {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, uuid)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `update ` + m.table + ` set ` + ` email=?, confirmed_and_active=1 ` + ` where id= ? limit = 1 `
		return conn.Exec(query, email, uuid)
	}, accountIdKey)
	return ret, err
}

func (m *AccountModel) UpdatePassword(uuid, passwordHash, salt string) (sql.Result, error) {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, uuid)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `update ` + m.table + ` set ` + ` password_hash=?, password_salt=? ` + ` where id=? limit=1 `
		return conn.Exec(query, passwordHash, salt, uuid)
	}, accountIdKey)
	return ret, err
}

func (m *AccountModel) Delete(id string) error {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `delete from ` + m.table + ` where id = ?`
		return conn.Exec(query, id)
	}, accountIdKey)
	return err
}

func (m *AccountModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAccountIdPrefix, primary)
}

func (m *AccountModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := `select ` + accountRows + ` from ` + m.table + ` where id = ? limit 1`
	return conn.QueryRow(v, query, primary)
}
