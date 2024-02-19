// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	lowProductListFieldNames          = builder.RawFieldNames(&LowProductList{})
	lowProductListRows                = strings.Join(lowProductListFieldNames, ",")
	lowProductListRowsExpectAutoSet   = strings.Join(stringx.Remove(lowProductListFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	lowProductListRowsWithPlaceHolder = strings.Join(stringx.Remove(lowProductListFieldNames, "`product_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	lowProductListModel interface {
		Insert(ctx context.Context, data *LowProductList) (sql.Result, error)
		FindOne(ctx context.Context, productId string) (*LowProductList, error)
		Update(ctx context.Context, data *LowProductList) error
		Delete(ctx context.Context, productId string) error
	}

	defaultLowProductListModel struct {
		conn  sqlx.SqlConn
		table string
	}

	LowProductList struct {
		ProductId       string    `db:"product_id"`       // 商品id
		ProductName     string    `db:"product_name"`     // 商品名称
		ProductTitle    string    `db:"product_title"`    // 商品描述
		ProductTypeId   string    `db:"product_type_id"`  // 商品类型id
		ProductQuantity int64     `db:"product_quantity"` // 商品数量
		ProductPicture  string    `db:"product_picture"`  // 商品图片
		Price           string    `db:"price"`            // 商品价格
		Producer        string    `db:"producer"`         // 产地
		Quota           int64     `db:"quota"`            // 限购数量
		ProductSize     int64     `db:"product_size"`     // 商品规格
		ShopId          string    `db:"shop_id"`          // 店铺id
		DeleteKey       int64     `db:"delete_key"`       // 删除虚拟key
		CreationTime    time.Time `db:"creation_time"`    // 创建时间
		UpdataTime      time.Time `db:"updata_time"`      // 修改时间
		StartTime       time.Time `db:"start_time"`       // 开始时间
		EndTime         time.Time `db:"end_time"`         // 截至时间
	}
)

func newLowProductListModel(conn sqlx.SqlConn) *defaultLowProductListModel {
	return &defaultLowProductListModel{
		conn:  conn,
		table: "`low_product_list`",
	}
}

func (m *defaultLowProductListModel) Delete(ctx context.Context, productId string) error {
	query := fmt.Sprintf("delete from %s where `product_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, productId)
	return err
}

func (m *defaultLowProductListModel) FindOne(ctx context.Context, productId string) (*LowProductList, error) {
	query := fmt.Sprintf("select %s from %s where `product_id` = ? limit 1", lowProductListRows, m.table)
	var resp LowProductList
	err := m.conn.QueryRowCtx(ctx, &resp, query, productId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLowProductListModel) Insert(ctx context.Context, data *LowProductList) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, lowProductListRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductId, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.Producer, data.Quota, data.ProductSize, data.ShopId, data.DeleteKey, data.CreationTime, data.UpdataTime, data.StartTime, data.EndTime)
	return ret, err
}

func (m *defaultLowProductListModel) Update(ctx context.Context, data *LowProductList) error {
	query := fmt.Sprintf("update %s set %s where `product_id` = ?", m.table, lowProductListRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.Producer, data.Quota, data.ProductSize, data.ShopId, data.DeleteKey, data.CreationTime, data.UpdataTime, data.StartTime, data.EndTime, data.ProductId)
	return err
}

func (m *defaultLowProductListModel) tableName() string {
	return m.table
}
