package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		TransactInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		PhoneSelectUser(ctx context.Context, session sqlx.Session, phone string) (*User, error)
		TransactSelectUserList(ctx context.Context, session sqlx.Session, limit int64) (*[]User, error)
		SelectUserListCount(ctx context.Context) (*UserListCount, error)
		SelectUserNumberTheDay(ctx context.Context) (*TheUserNumber, error)
	}

	customUserModel struct {
		*defaultUserModel
	}

	UserListCount struct {
		UserCount int64 `db:"user_count"`
	}

	TheUserNumber struct {
		UserCount int64 `db:"user_count"`
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (c customUserModel) TransactInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", c.table, userRowsExpectAutoSet)
	if session == nil {
		ret, err := c.conn.ExecCtx(ctx, query, data.Id, data.Name, data.Phone, data.Password, data.PositionId, data.RegistrationTime)
		return ret, err
	} else {
		ret, err := session.ExecCtx(ctx, query, data.Id, data.Name, data.Phone, data.Password, data.PositionId, data.RegistrationTime)
		return ret, err
	}
}

func (c customUserModel) PhoneSelectUser(ctx context.Context, session sqlx.Session, phone string) (*User, error) {
	query := fmt.Sprintf("select * from %s where `phone` = ? limit 1", c.table)
	var resp User
	if session == nil {
		err := c.conn.QueryRowCtx(ctx, &resp, query, phone)
		if err != nil {
			return nil, err
		}
	} else {
		err := session.QueryRowCtx(ctx, &resp, query, phone)
		if err != nil {
			return nil, err
		}
	}
	return &resp, nil
}

func (c *customUserModel) TransactSelectUserList(ctx context.Context, session sqlx.Session, limit int64) (*[]User, error) {
	query := fmt.Sprintf("select * from %s limit ? , 10", c.table)
	var resp []User
	if session == nil {
		err := c.conn.QueryRowsCtx(ctx, &resp, query, limit)
		if err != nil {
			return nil, err
		}
	} else {
		err := session.QueryRowsCtx(ctx, &resp, query, limit)
		if err != nil {
			return nil, err
		}
	}
	return &resp, nil
}

func (c *customUserModel) SelectUserListCount(ctx context.Context) (*UserListCount, error) {
	query := fmt.Sprintf("select COUNT(*) AS `user_count` from %s ; ", c.table)
	var resp UserListCount
	err := c.conn.QueryRowCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *customUserModel) SelectUserNumberTheDay(ctx context.Context) (*TheUserNumber, error) {
	query := fmt.Sprintf("select COUNT(*) AS `user_count` from %s  where DATE(`registration_time`)=CURDATE(); ", c.table)
	var resp TheUserNumber
	err := c.conn.QueryRowCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
