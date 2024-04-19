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

var _ ProductListModel = (*customProductListModel)(nil)

type (
	// ProductListModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductListModel.
	ProductListModel interface {
		productListModel
		SelectAllLimit(ctx context.Context, in *sms.ProductListReq) (*[]ProductALL, error)
		SelectProductListLimit(ctx context.Context, in *sms.ProductListReq) (*[]ProductALL, error)
		TransactSelectProductList(ctx context.Context, session sqlx.Session, id string) (*ProductList, error)
		TransactUpDataProductList(ctx context.Context, session sqlx.Session, id string, number int) error
		SelectProductListByShopId(ctx context.Context, session sqlx.Session, shopId string, limit int) (*[]ProductList, error)
		TransactInsert(ctx context.Context, session sqlx.Session, data *ProductList) (sql.Result, error)
		TransactUpdateProductData(ctx context.Context, session sqlx.Session, data *ProductList) error
		SelectProductSumByShopId(ctx context.Context, shopId string) (*ProductSum, error)
	}

	customProductListModel struct {
		*defaultProductListModel
	}

	ProductALL struct {
		ProductId       string `db:"product_id"`        // 商品id
		ProductName     string `db:"product_name"`      // 商品名称
		ProductPicture  string `db:"product_picture"`   // 商品图片
		ProductTypeName string `db:"product_type_name"` // 商品类型名称
		Price           string `db:"price"`             // 商品价格
		ProductSize     int64  `db:"product_size"`      // 商品规格
		Producer        string `db:"producer"`          // 产地
	}

	ProductSum struct {
		ShopId      string `db:"shop_id"`      // 商品id
		CountNumber int64  `db:"count_number"` // 统计商品数量
	}
)

// NewProductListModel returns a model for the database table.
func NewProductListModel(conn sqlx.SqlConn) ProductListModel {
	return &customProductListModel{
		defaultProductListModel: newProductListModel(conn),
	}
}

func (c customProductListModel) SelectAllLimit(ctx context.Context, in *sms.ProductListReq) (*[]ProductALL, error) {
	query := "SELECT product_id,product_name,product_picture,t.`product_type_name`,`price`,`product_size`,`producer`FROM product_list l JOIN product_type t ON l.product_type_id=t.product_type_id WHERE shop_id = ? AND delete_key=0 AND product_quantity>0 LIMIT ? , ?"

	var resp []ProductALL
	err := c.conn.QueryRowsCtx(ctx, &resp, query, in.ShopId, in.Quantity, in.Limit)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c customProductListModel) SelectProductListLimit(ctx context.Context, in *sms.ProductListReq) (*[]ProductALL, error) {
	query := "SELECT product_id,product_name,product_picture,t.`product_type_name`,`price`,`product_size`,`producer`FROM product_list l JOIN product_type t ON l.product_type_id=t.product_type_id WHERE shop_id = ? AND delete_key=0 AND product_quantity>0 AND T.`product_type_name`= ? LIMIT ? , ?"
	var resp []ProductALL
	err := c.conn.QueryRowsCtx(ctx, &resp, query, in.ShopId, in.ProductType, in.Quantity, in.Limit)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *customProductListModel) TransactSelectProductList(ctx context.Context, session sqlx.Session, id string) (*ProductList, error) {

	query := fmt.Sprintf("select %s from %s where `product_id` = ? limit 1 for update", productListRows, c.table)
	var resp ProductList
	var err error
	if session == nil {
		err = c.conn.QueryRowCtx(ctx, &resp, query, id)
	} else {
		err = session.QueryRowCtx(ctx, &resp, query, id)
	}

	return &resp, err
}

func (c *customProductListModel) TransactUpDataProductList(ctx context.Context, session sqlx.Session, id string, number int) error {
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

func (c *customProductListModel) SelectProductListByShopId(ctx context.Context, session sqlx.Session, shopId string, limit int) (*[]ProductList, error) {

	query := fmt.Sprintf("select %s from %s where `shop_id` = ? and delete_key = 0 limit ? , 10", productListRows, c.table)

	var resp []ProductList
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

func (c *customProductListModel) TransactInsert(ctx context.Context, session sqlx.Session, data *ProductList) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", c.table, productListRowsExpectAutoSet)
	if session != nil {
		ret, err := session.ExecCtx(ctx, query, data.ProductId, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.ProductSize, data.ShopId, data.Producer, data.DeleteKey, data.CreationTime, data.UpdataTime)
		return ret, err
	} else {
		ret, err := c.conn.ExecCtx(ctx, query, data.ProductId, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.ProductSize, data.ShopId, data.Producer, data.DeleteKey, data.CreationTime, data.UpdataTime)
		return ret, err
	}
}

func (c *customProductListModel) TransactUpdateProductData(ctx context.Context, session sqlx.Session, data *ProductList) error {
	query := fmt.Sprintf("update %s set %s where `product_id` = ?", c.table, productListRowsWithPlaceHolder)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.ProductSize, data.ShopId, data.Producer, data.DeleteKey, data.CreationTime, data.UpdataTime, data.ProductId)
		return err
	} else {
		_, err := c.conn.ExecCtx(ctx, query, data.ProductName, data.ProductTitle, data.ProductTypeId, data.ProductQuantity, data.ProductPicture, data.Price, data.ProductSize, data.ShopId, data.Producer, data.DeleteKey, data.CreationTime, data.UpdataTime, data.ProductId)
		return err
	}
}
func (c *customProductListModel) SelectProductSumByShopId(ctx context.Context, shopId string) (*ProductSum, error) {
	query := fmt.Sprintf("SELECT `shop_id` , COUNT(*) AS `count_number` FROM %s WHERE shop_id = ? GROUP BY `shop_id`;", c.table)
	var resp ProductSum
	err := c.conn.QueryRowCtx(ctx, &resp, query, shopId)
	return &resp, err
}
