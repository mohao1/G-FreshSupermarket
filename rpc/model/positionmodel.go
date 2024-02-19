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
