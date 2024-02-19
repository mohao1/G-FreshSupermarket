package model

import (
	"DP/rpc/Sms/pb/sms"
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"time"
)

var _ LowProductListModel = (*customLowProductListModel)(nil)

type (
	// LowProductListModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLowProductListModel.
	LowProductListModel interface {
		lowProductListModel
		SelectAllLimit(ctx context.Context, in *sms.LowProductListReq) (*[]LowProductALL, error)
		SelectLowProductListLimit(ctx context.Context, in *sms.LowProductListReq) (*[]LowProductALL, error)
		TransactSelectLowProductList(ctx context.Context, session sqlx.Session, id string) (*LowProductList, error)
		TransactUpDataLowProductList(ctx context.Context, session sqlx.Session, id string, number int) error
		SelectLowProductListByShopId(ctx context.Context, session sqlx.Session, shopId string, limit int64) (*[]LowProductList, error)
		TransactUpdateLowProductData(ctx context.Context, session sqlx.Session, data *LowProductList) error
		TransactInsert(ctx context.Context, session sqlx.Session, data *LowProductList) (sql.Result, error)
	}

	customLowProductListModel struct {
		*defaultLowProductListModel
	}

	LowProductALL struct {
		ProductId       string    `db:"product_id"`        // 商品id
		ProductName     string    `db:"product_name"`      // 商品名称
		ProductPicture  string    `db:"product_picture"`   // 商品图片
		ProductTypeName string    `db:"product_type_name"` // 商品类型名称
		Price           string    `db:"price"`             // 商品价格
		ProductSize     int64     `db:"product_size"`      // 商品规格
		Producer        string    `db:"producer"`          // 产地
		Quota           int64     `db:"quota"`             // 限购数量
		StartTime       time.Time `db:"start_time"`        // 开始时间
		EndTime         time.Time `db:"end_time"`          // 截至时间

	}
)

// NewLowProductListModel returns a model for the database table.
func NewLowProductListModel(conn sqlx.SqlConn) LowProductListModel {
	return &customLowProductListModel{
		defaultLowProductListModel: newLowProductListModel(conn),
	}
}

func (c *customLowProductListModel) SelectAllLimit(ctx context.Context, in *sms.LowProductListReq) (*[]LowProductALL, error) {
	//query := "SELECT product_id,product_name,product_picture,t.`product_type_name`,`price`,`product_size`,`producer`FROM low_product_list l JOIN product_type t ON l.product_type_id=t.product_type_id WHERE shop_id = ? AND delete_key=0 AND product_quantity>0 LIMIT ? , ?"
	Rows := strings.Join(builder.RawFieldNames(&LowProductALL{}), ",")
	query := fmt.Sprintf("SELECT %s FROM low_product_list l JOIN product_type t ON l.product_type_id=t.product_type_id WHERE shop_id = ? AND delete_key=0 AND product_quantity>0 AND `start_time` < NOW() AND `end_time` > NOW() LIMIT ? , ?", Rows)
	var resp []LowProductALL
	err := c.conn.QueryRowsCtx(ctx, &resp, query, in.ShopId, in.Quantity, in.Limit)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *customLowProductListModel) SelectLowProductListLimit(ctx context.Context, in *sms.LowProductListReq) (*[]LowProductALL, error) {
	Rows := strings.Join(builder.RawFieldNames(&LowProductALL{}), ",")
	query := fmt.Sprintf("SELECT %s FROM low_product_list l JOIN product_type t ON l.product_type_id=t.product_type_id WHERE shop_id = ? AND delete_key=0 AND product_quantity>0 AND product_type_name=? AND `start_time` < NOW() AND `end_time` > NOW() LIMIT ? , ?", Rows)
	var resp []LowProductALL
	err := c.conn.QueryRowsCtx(ctx, &resp, query, in.ShopId, in.ProductType, in.Quantity, in.Limit)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *customLowProductListModel) TransactSelectLowProductList(ctx context.Context, session sqlx.Session, id string) (*LowProductList, error) {
	query := fmt.Sprintf("select %s from %s where `product_id` = ? limit 1 for update", lowProductListRows, c.table)
	var resp LowProductList
	var err error
	if session == nil {
		err = c.conn.QueryRowCtx(ctx, &resp, query, id)
	} else {
		err = session.QueryRowCtx(ctx, &resp, query, id)
	}

	return &resp, err
}

func (c *customLowProductListModel) TransactUpDataLowProductList(ctx context.Context, session sqlx.Session, id string, number int) error {
	type Req struct {
		ProductQuantity int64     `db:"product_quantity"` // 商品数量
		UpdataTime      time.Time `db:"updata_time"`      // 修改时间
	}
	Rows := strings.Join(builder.RawFieldNames(&Req{}), "=?,") + "=?"
	query := fmt.Sprintf("update %s set %s where `product_id` = ?", c.table, Rows)
	var err error
	if session == nil {
		_, err = c.conn.ExecCtx(ctx, query, number, time.Now(), id)

	} else {
		_, err = session.ExecCtx(ctx, query, number, time.Now(), id)
	}
	return err
}

func (c *customLowProductListModel) SelectLowProductListByShopId(ctx context.Context, session sqlx.Session, shopId string, limit int64) (*[]LowProductList, error) {
	query := fmt.Sprintf("select %s from %s where `shop_id` = ? and `delete_key` = 0 limit ? , 10", lowProductListRows, c.table)

	var resp []LowProductList
	if session != nil {
		err := session.QueryRowsCtx(ctx, &resp, query, shopId, limit)
		if err != nil {
			return nil, err
		}
	} else {
		err := c.conn.QueryRowsCtx(ctx, &resp, query, shopId, limit)
		if err != nil {
			return nil, err
		}
	}

	return &resp, nil
}

func (c *customLowProductListModel) TransactUpdateLowProductData(ctx context.Context, session sqlx.Session, data *LowProductList) error {
	query := fmt.Sprintf("update %s set %s where `product_id` = ?", c.table, lowProductListRowsWithPlaceHolder)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.Producer, data.Quota, data.ProductSize, data.ShopId, data.DeleteKey, data.CreationTime, data.UpdataTime, data.StartTime, data.EndTime, data.ProductId)
		return err
	} else {
		_, err := c.conn.ExecCtx(ctx, query, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.Producer, data.Quota, data.ProductSize, data.ShopId, data.DeleteKey, data.CreationTime, data.UpdataTime, data.StartTime, data.EndTime, data.ProductId)
		return err
	}
}

func (c *customLowProductListModel) TransactInsert(ctx context.Context, session sqlx.Session, data *LowProductList) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", c.table, lowProductListRowsExpectAutoSet)
	if session != nil {
		ret, err := session.ExecCtx(ctx, query, data.ProductId, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.Producer, data.Quota, data.ProductSize, data.ShopId, data.DeleteKey, data.CreationTime, data.UpdataTime, data.StartTime, data.EndTime)
		return ret, err
	} else {
		ret, err := c.conn.ExecCtx(ctx, query, data.ProductId, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.Producer, data.Quota, data.ProductSize, data.ShopId, data.DeleteKey, data.CreationTime, data.UpdataTime, data.StartTime, data.EndTime)
		return ret, err
	}
}
