package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PositionModel = (*customPositionModel)(nil)

type (
	// PositionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionModel.
	PositionModel interface {
		positionModel
		SelectPositionList(ctx context.Context) (*[]Position, error)
		TransactSelectPosition(ctx context.Context, session sqlx.Session, positionId string) (*Position, error)
		TransactUpDatePosition(ctx context.Context, session sqlx.Session, data *Position) error
		TransactDeletePosition(ctx context.Context, session sqlx.Session, positionId string) error
	}

	customPositionModel struct {
		*defaultPositionModel
	}
)

// NewPositionModel returns a model for the database table.
func NewPositionModel(conn sqlx.SqlConn) PositionModel {
	return &customPositionModel{
		defaultPositionModel: newPositionModel(conn),
	}
}

func (c *customPositionModel) SelectPositionList(ctx context.Context) (*[]Position, error) {
	query := fmt.Sprintf("select %s from %s ", positionRows, c.table)
	var resp []Position
	err := c.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *customPositionModel) TransactSelectPosition(ctx context.Context, session sqlx.Session, positionId string) (*Position, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1 for update", positionRows, c.table)
	var resp Position
	if session != nil {
		err := session.QueryRowCtx(ctx, &resp, query, positionId)
		return &resp, err
	} else {
		err := c.conn.QueryRowCtx(ctx, &resp, query, positionId)
		return &resp, err
	}
}

func (c *customPositionModel) TransactUpDatePosition(ctx context.Context, session sqlx.Session, data *Position) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", c.table, positionRowsWithPlaceHolder)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, data.PositionName, data.Grade, data.Id)
		return err
	} else {
		_, err := c.conn.ExecCtx(ctx, query, data.PositionName, data.Grade, data.Id)
		return err
	}
}

func (c *customPositionModel) TransactDeletePosition(ctx context.Context, session sqlx.Session, positionId string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", c.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, positionId)
		return err
	} else {
		_, err := c.conn.ExecCtx(ctx, query, positionId)
		return err
	}
}
