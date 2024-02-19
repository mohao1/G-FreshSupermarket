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
	}

	customUserModel struct {
		*defaultUserModel
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
