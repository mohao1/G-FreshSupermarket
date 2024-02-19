package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductTypeModel = (*customProductTypeModel)(nil)

type (
	// ProductTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductTypeModel.
	ProductTypeModel interface {
		productTypeModel
		SelectType(ctx context.Context, typeStr string) (*ProductType, error)
		SelectTypeList(ctx context.Context) (*[]ProductType, error)
	}

	customProductTypeModel struct {
		*defaultProductTypeModel
	}
)

// NewProductTypeModel returns a model for the database table.
func NewProductTypeModel(conn sqlx.SqlConn) ProductTypeModel {
	return &customProductTypeModel{
		defaultProductTypeModel: newProductTypeModel(conn),
	}
}

func (c *customProductTypeModel) SelectType(ctx context.Context, typeStr string) (*ProductType, error) {
	query := fmt.Sprintf("select * from %s where `product_type_name` = ? limit 1", c.table)
	var resp ProductType
	err := c.conn.QueryRowCtx(ctx, &resp, query, typeStr)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *customProductTypeModel) SelectTypeList(ctx context.Context) (*[]ProductType, error) {
	query := fmt.Sprintf("select * from %s", c.table)
	var resp []ProductType
	err := c.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
