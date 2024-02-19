package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SalesRecordsModel = (*customSalesRecordsModel)(nil)

type (
	// SalesRecordsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSalesRecordsModel.
	SalesRecordsModel interface {
		salesRecordsModel
		SelectSalesRecordsListByShopId(ctx context.Context, shopId string) (*[]SalesRecords, error)
		TransactInsert(ctx context.Context, session sqlx.Session, data *SalesRecords) (sql.Result, error)
	}

	customSalesRecordsModel struct {
		*defaultSalesRecordsModel
	}
)

// NewSalesRecordsModel returns a model for the database table.
func NewSalesRecordsModel(conn sqlx.SqlConn) SalesRecordsModel {
	return &customSalesRecordsModel{
		defaultSalesRecordsModel: newSalesRecordsModel(conn),
	}
}

func (c *customSalesRecordsModel) SelectSalesRecordsListByShopId(ctx context.Context, shopId string) (*[]SalesRecords, error) {
	query := fmt.Sprintf("select %s from %s where `shop_id` = ?", salesRecordsRows, c.table)
	var resp []SalesRecords
	err := c.conn.QueryRowsCtx(ctx, &resp, query, shopId)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *customSalesRecordsModel) TransactInsert(ctx context.Context, session sqlx.Session, data *SalesRecords) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", c.table, salesRecordsRowsExpectAutoSet)
	if session != nil {
		ret, err := session.ExecCtx(ctx, query, data.SalesRecordsId, data.ProductId, data.SalesQuantity, data.TotalPrice, data.ShopId, data.CreationTime, data.UpdataTime)
		return ret, err
	} else {
		ret, err := c.conn.ExecCtx(ctx, query, data.SalesRecordsId, data.ProductId, data.SalesQuantity, data.TotalPrice, data.ShopId, data.CreationTime, data.UpdataTime)
		return ret, err
	}
}
