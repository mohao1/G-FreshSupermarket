package model

import (
	"context"
	"database/sql"
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
		SelectShopList(ctx context.Context, limit int64) (*[]Shop, error)
		TransactInsertShop(ctx context.Context, session sqlx.Session, data *Shop) (sql.Result, error)
		TransactSelectShop(ctx context.Context, session sqlx.Session, shopId string) (*Shop, error)
		TransactUpDateShop(ctx context.Context, session sqlx.Session, data *Shop) error
		TransactDeleteShop(ctx context.Context, session sqlx.Session, shopId string) error
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

func (m *customShopModel) SelectShopList(ctx context.Context, limit int64) (*[]Shop, error) {
	query := fmt.Sprintf("select %s from %s limit ? , 10", shopRows, m.table)
	var resp []Shop
	err := m.conn.QueryRowsCtx(ctx, &resp, query, limit)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *customShopModel) TransactInsertShop(ctx context.Context, session sqlx.Session, data *Shop) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ? , ?)", m.table, shopRowsExpectAutoSet)
	if session != nil {
		ret, err := session.ExecCtx(ctx, query, data.ShopId, data.ShopName, data.ShopAddress, data.ShopCity, data.ShopAdmin, data.CreationTime)
		return ret, err
	} else {
		ret, err := m.conn.ExecCtx(ctx, query, data.ShopId, data.ShopName, data.ShopAddress, data.ShopCity, data.ShopAdmin, data.CreationTime)
		return ret, err
	}
}

func (m *customShopModel) TransactSelectShop(ctx context.Context, session sqlx.Session, shopId string) (*Shop, error) {
	query := fmt.Sprintf("select %s from %s where `shop_id` = ? limit 1 for update", shopRows, m.table)
	var resp Shop
	if session != nil {
		err := session.QueryRowCtx(ctx, &resp, query, shopId)
		return &resp, err
	} else {
		err := m.conn.QueryRowCtx(ctx, &resp, query, shopId)
		return &resp, err
	}
}

func (m *customShopModel) TransactUpDateShop(ctx context.Context, session sqlx.Session, data *Shop) error {
	query := fmt.Sprintf("update %s set %s where `shop_id` = ?", m.table, shopRowsWithPlaceHolder)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, data.ShopName, data.ShopAddress, data.ShopCity, data.ShopAdmin, data.CreationTime, data.ShopId)
		return err
	} else {
		_, err := m.conn.ExecCtx(ctx, query, data.ShopName, data.ShopAddress, data.ShopCity, data.ShopAdmin, data.CreationTime, data.ShopId)
		return err
	}

}

func (m *customShopModel) TransactDeleteShop(ctx context.Context, session sqlx.Session, shopId string) error {

	query := fmt.Sprintf("delete from %s where `shop_id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, shopId)
		return err
	} else {
		_, err := m.conn.ExecCtx(ctx, query, shopId)
		return err
	}
}
