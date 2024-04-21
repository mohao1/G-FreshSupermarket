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
		SelectSalesRecordsListByShopId(ctx context.Context, ShopId string) (*[]SalesRecordsData, error)
		SelectShopSalesRecordsSum(ctx context.Context, ShopId string) (*ShopSalesRecordsSum, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}

	SalesRecordsData struct {
		ProductId              string `db:"product_id"`                // 商品id
		ProductPrice           string `db:"product_price"`             // 价格
		ProductSalesRecordsSum int64  `db:"product_sales_records_sum"` //统计数量
	}

	ShopSalesRecordsSum struct {
		ProductSalesRecordsSum int64 `db:"product_sales_records_sum"` //统计数量
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn),
	}
}

func (c *customOrderModel) TransactInsert(ctx context.Context, session sqlx.Session, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", c.table, orderRowsExpectAutoSet)
	if session == nil {
		ret, err := c.conn.ExecCtx(ctx, query, data.OrderId, data.OrderName, data.OrderTitle, data.Price, data.ProductId, data.ProductTypeId, data.OrderQuantity, data.ProductSize, data.ProductPicture, data.OrderNumber, data.ShopId, data.CreationTime, data.UpdataTime, data.DeleteKey)
		return ret, err
	} else {
		ret, err := session.ExecCtx(ctx, query, data.OrderId, data.OrderName, data.OrderTitle, data.Price, data.ProductId, data.ProductTypeId, data.OrderQuantity, data.ProductSize, data.ProductPicture, data.OrderNumber, data.ShopId, data.CreationTime, data.UpdataTime, data.DeleteKey)
		return ret, err
	}
}

func (c *customOrderModel) SelectOrderByOrderNumberId(ctx context.Context, orderNumber string) (*[]Order, error) {
	query := fmt.Sprintf("select %s from %s where `order_number` = ?", orderRows, c.table)
	var resp []Order
	err := c.conn.QueryRowsCtx(ctx, &resp, query, orderNumber)
	return &resp, err
}

func (c *customOrderModel) SelectSalesRecordsListByShopId(ctx context.Context, ShopId string) (*[]SalesRecordsData, error) {
	query := fmt.Sprintf("select `product_id` , SUM(price)  AS `product_price`, SUM(order_quantity) AS `product_sales_records_sum` from %s where `shop_id` = ? GROUP BY `product_id`", c.table)
	var resp []SalesRecordsData
	err := c.conn.QueryRowsCtx(ctx, &resp, query, ShopId)
	return &resp, err
}

func (c *customOrderModel) SelectShopSalesRecordsSum(ctx context.Context, ShopId string) (*ShopSalesRecordsSum, error) {
	query := fmt.Sprintf("select SUM(order_quantity) AS `product_sales_records_sum` from %s where `shop_id` = ?;", c.table)
	var resp ShopSalesRecordsSum
	err := c.conn.QueryRowCtx(ctx, &resp, query, ShopId)
	return &resp, err
}
