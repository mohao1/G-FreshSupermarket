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
		TransactSelectProductType(ctx context.Context, session sqlx.Session, productTypeId string) (*ProductType, error)
		TransactUpDateProductType(ctx context.Context, session sqlx.Session, data *ProductType) error
		TransactDeleteProductType(ctx context.Context, session sqlx.Session, productTypeId string) error
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

func (c *customProductTypeModel) TransactSelectProductType(ctx context.Context, session sqlx.Session, productTypeId string) (*ProductType, error) {
	query := fmt.Sprintf("select %s from %s where `product_type_id` = ? limit 1 for update", productTypeRows, c.table)
	if session != nil {
		var resp ProductType
		err := session.QueryRowCtx(ctx, &resp, query, productTypeId)
		return &resp, err
	} else {
		var resp ProductType
		err := c.conn.QueryRowCtx(ctx, &resp, query, productTypeId)
		return &resp, err
	}

}

func (c *customProductTypeModel) TransactUpDateProductType(ctx context.Context, session sqlx.Session, data *ProductType) error {
	query := fmt.Sprintf("update %s set %s where `product_type_id` = ?", c.table, productTypeRowsWithPlaceHolder)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, data.ProductTypeName, data.ProductUnit, data.ProductTypeId)
		return err
	} else {
		_, err := c.conn.ExecCtx(ctx, query, data.ProductTypeName, data.ProductUnit, data.ProductTypeId)
		return err
	}
}

func (c *customProductTypeModel) TransactDeleteProductType(ctx context.Context, session sqlx.Session, productTypeId string) error {
	query := fmt.Sprintf("delete from %s where `product_type_id` = ?", c.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, productTypeId)
		return err
	} else {
		_, err := c.conn.ExecCtx(ctx, query, productTypeId)
		return err
	}
}
