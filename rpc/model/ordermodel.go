package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		TransactInsert(ctx context.Context, session sqlx.Session, order *Order) (sql.Result, error)
		SelectOrderByOrderNumberId(ctx context.Context, orderNumber string) (*[]Order, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn),
	}
}

func (c *customOrderModel) TransactInsert(ctx context.Context, session sqlx.Session, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", c.table, orderRowsExpectAutoSet)
	if session == nil {
		ret, err := c.conn.ExecCtx(ctx, query, data.OrderId, data.OrderName, data.OrderTitle, data.Price, data.ProductTypeId, data.OrderQuantity, data.ProductSize, data.ProductPicture, data.OrderNumber, data.ShopId, data.CreationTime, data.UpdataTime, data.DeleteKey)
		return ret, err
	} else {
		ret, err := session.ExecCtx(ctx, query, data.OrderId, data.OrderName, data.OrderTitle, data.Price, data.ProductTypeId, data.OrderQuantity, data.ProductSize, data.ProductPicture, data.OrderNumber, data.ShopId, data.CreationTime, data.UpdataTime, data.DeleteKey)
		return ret, err
	}
}

func (c *customOrderModel) SelectOrderByOrderNumberId(ctx context.Context, orderNumber string) (*[]Order, error) {
	query := fmt.Sprintf("select %s from %s where `order_number` = ?", orderRows, c.table)
	var resp []Order
	err := c.conn.QueryRowsCtx(ctx, &resp, query, orderNumber)
	return &resp, err
}
