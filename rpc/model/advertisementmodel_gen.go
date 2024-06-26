// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	advertisementFieldNames          = builder.RawFieldNames(&Advertisement{})
	advertisementRows                = strings.Join(advertisementFieldNames, ",")
	advertisementRowsExpectAutoSet   = strings.Join(stringx.Remove(advertisementFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	advertisementRowsWithPlaceHolder = strings.Join(stringx.Remove(advertisementFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	advertisementModel interface {
		Insert(ctx context.Context, data *Advertisement) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Advertisement, error)
		Update(ctx context.Context, data *Advertisement) error
		Delete(ctx context.Context, id string) error
	}

	defaultAdvertisementModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Advertisement struct {
		Id      string `db:"id"`      // 广告id
		Picture string `db:"picture"` // 轮播图片
		ShopId  string `db:"shop_id"` // 店铺id
	}
)

func newAdvertisementModel(conn sqlx.SqlConn) *defaultAdvertisementModel {
	return &defaultAdvertisementModel{
		conn:  conn,
		table: "`advertisement`",
	}
}

func (m *defaultAdvertisementModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAdvertisementModel) FindOne(ctx context.Context, id string) (*Advertisement, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", advertisementRows, m.table)
	var resp Advertisement
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdvertisementModel) Insert(ctx context.Context, data *Advertisement) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, advertisementRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Picture, data.ShopId)
	return ret, err
}

func (m *defaultAdvertisementModel) Update(ctx context.Context, data *Advertisement) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, advertisementRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Picture, data.ShopId, data.Id)
	return err
}

func (m *defaultAdvertisementModel) tableName() string {
	return m.table
}
