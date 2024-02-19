package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShopModel = (*customShopModel)(nil)

type (
	// ShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShopModel.
	ShopModel interface {
		shopModel
		SelectShopByCity(ctx context.Context, city string) (*[]Shop, error)
	}

	customShopModel struct {
		*defaultShopModel
	}
)

// NewShopModel returns a model for the database table.
func NewShopModel(conn sqlx.SqlConn) ShopModel {
	return &customShopModel{
		defaultShopModel: newShopModel(conn),
	}
}

func (m *customShopModel) SelectShopByCity(ctx context.Context, city string) (*[]Shop, error) {
	query := fmt.Sprintf("select %s from %s where `shop_city` = ?", shopRows, m.table)
	var resp []Shop
	err := m.conn.QueryRowsCtx(ctx, &resp, query, city)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
