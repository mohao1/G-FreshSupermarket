package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"time"
)

var _ OrderNumberModel = (*customOrderNumberModel)(nil)

type (
	// OrderNumberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderNumberModel.
	OrderNumberModel interface {
		orderNumberModel
		TransactInsert(ctx context.Context, session sqlx.Session, data *OrderNumber) (sql.Result, error)
		SelectOrderNumber(ctx context.Context, userId string, limit string) (*[]OrderNumber, error)
		SelectOrderNumberByUserIdAndOrderNumberId(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) (*OrderNumber, error)
		UpDataOrderConfirm(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) error
		UpDataOrderOver(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) error
		UpDataOrderRefund(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) error
		UpDataOrderUnRefund(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) error
		SelectReceivedOrderListByShop(ctx context.Context, session sqlx.Session, shopId string, receive int, confirmedDeliver int, limit int) (*[]OrderNumber, error)
		SelectOrderNumberByShopIdAndOrderNumberId(ctx context.Context, session sqlx.Session, shopId string, orderNumberId string) (*OrderNumber, error)
		SelectOrderDataListByShopAndTime(ctx context.Context, shopId string, time time.Time) (*[]OrderNumber, error)
		SelectOrderNumberByOrderId(ctx context.Context, session sqlx.Session, orderId string) (*OrderNumber, error)
		UpDateOrderReceive(ctx context.Context, session sqlx.Session, receive int64, orderNumberId string) error
	}

	customOrderNumberModel struct {
		*defaultOrderNumberModel
	}
)

// NewOrderNumberModel returns a model for the database table.
func NewOrderNumberModel(conn sqlx.SqlConn) OrderNumberModel {
	return &customOrderNumberModel{
		defaultOrderNumberModel: newOrderNumberModel(conn),
	}
}

func (o customOrderNumberModel) TransactInsert(ctx context.Context, session sqlx.Session, data *OrderNumber) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", o.table, orderNumberRowsExpectAutoSet)

	if session == nil {
		ret, err := o.conn.ExecCtx(ctx, query, data.OrderNumber, data.CustomerId, data.TotalPrice, data.Total, data.Payment, data.ShopId, data.OrderOver, data.OrderReceive, data.ConfirmedDelivery, data.CreationTime, data.DeliveryTime, data.ConfirmTime, data.UpdataTime, data.EndTime, data.DeleteKey, data.Notes)
		return ret, err
	} else {
		ret, err := session.ExecCtx(ctx, query, data.OrderNumber, data.CustomerId, data.TotalPrice, data.Total, data.Payment, data.ShopId, data.OrderOver, data.OrderReceive, data.ConfirmedDelivery, data.CreationTime, data.DeliveryTime, data.ConfirmTime, data.UpdataTime, data.EndTime, data.DeleteKey, data.Notes)
		return ret, err
	}

}

func (o *customOrderNumberModel) SelectOrderNumber(ctx context.Context, userId string, limit string) (*[]OrderNumber, error) {
	query := fmt.Sprintf("select %s from %s where `customer_id` = ? limit ? , 8", orderNumberRows, o.table)
	var resp []OrderNumber
	err := o.conn.QueryRowsCtx(ctx, &resp, query, userId, limit)
	return &resp, err
}

func (o *customOrderNumberModel) SelectOrderNumberByUserIdAndOrderNumberId(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) (*OrderNumber, error) {
	query := fmt.Sprintf("select %s from %s where `customer_id` = ? And order_number = ? limit 1 for update", orderNumberRows, o.table)
	var resp OrderNumber

	if session != nil {
		err := session.QueryRowCtx(ctx, &resp, query, userId, orderNumberId)
		return &resp, err
	} else {
		err := o.conn.QueryRowCtx(ctx, &resp, query, userId, orderNumberId)
		return &resp, err
	}

}

func (o *customOrderNumberModel) SelectOrderNumberByShopIdAndOrderNumberId(ctx context.Context, session sqlx.Session, shopId string, orderNumberId string) (*OrderNumber, error) {
	query := fmt.Sprintf("select %s from %s where `shop_id` = ? And order_number = ? limit 1 for update", orderNumberRows, o.table)
	var resp OrderNumber

	if session != nil {
		err := session.QueryRowCtx(ctx, &resp, query, shopId, orderNumberId)
		return &resp, err
	} else {
		err := o.conn.QueryRowCtx(ctx, &resp, query, shopId, orderNumberId)
		return &resp, err
	}

}

func (o *customOrderNumberModel) UpDataOrderConfirm(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) error {

	type Req struct {
		ConfirmedDelivery int64        `db:"confirmed_delivery"` // 确认交付标记
		ConfirmTime       sql.NullTime `db:"confirm_time"`       // 确认时间
		UpdataTime        time.Time    `db:"updata_time"`        // 修改时间
	}
	Rows := strings.Join(builder.RawFieldNames(&Req{}), "=?,") + "=?"
	query := fmt.Sprintf("update %s set %s where `order_number` = ? AND `customer_id` = ?", o.table, Rows)
	var err error
	if session == nil {
		_, err = o.conn.ExecCtx(ctx, query, 1, time.Now(), time.Now(), orderNumberId, userId)

	} else {
		_, err = session.ExecCtx(ctx, query, 1, time.Now(), time.Now(), orderNumberId, userId)
	}
	return err

}

func (o *customOrderNumberModel) UpDataOrderOver(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) error {

	type Req struct {
		OrderOver  int64     `db:"order_over"`  // 是否已经取消
		UpdataTime time.Time `db:"updata_time"` // 修改时间
	}
	Rows := strings.Join(builder.RawFieldNames(&Req{}), "=?,") + "=?"
	query := fmt.Sprintf("update %s set %s where `order_number` = ? AND `customer_id` = ?", o.table, Rows)
	var err error
	if session == nil {
		_, err = o.conn.ExecCtx(ctx, query, 1, time.Now(), orderNumberId, userId)

	} else {
		_, err = session.ExecCtx(ctx, query, 1, time.Now(), orderNumberId, userId)
	}
	return err

}

func (o *customOrderNumberModel) UpDataOrderRefund(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) error {

	type Req struct {
		OrderOver  int64     `db:"order_over"`  // 是否已经取消
		UpdataTime time.Time `db:"updata_time"` // 修改时间
	}
	Rows := strings.Join(builder.RawFieldNames(&Req{}), "=?,") + "=?"
	query := fmt.Sprintf("update %s set %s where `order_number` = ? AND `customer_id` = ?", o.table, Rows)
	var err error
	if session == nil {
		_, err = o.conn.ExecCtx(ctx, query, 2, time.Now(), orderNumberId, userId)

	} else {
		_, err = session.ExecCtx(ctx, query, 2, time.Now(), orderNumberId, userId)
	}
	return err

}

func (o *customOrderNumberModel) UpDataOrderUnRefund(ctx context.Context, session sqlx.Session, userId string, orderNumberId string) error {

	type Req struct {
		OrderOver  int64     `db:"order_over"`  // 是否已经取消
		UpdataTime time.Time `db:"updata_time"` // 修改时间
	}
	Rows := strings.Join(builder.RawFieldNames(&Req{}), "=?,") + "=?"
	query := fmt.Sprintf("update %s set %s where `order_number` = ? AND `customer_id` = ?", o.table, Rows)
	var err error
	if session == nil {
		_, err = o.conn.ExecCtx(ctx, query, 0, time.Now(), orderNumberId, userId)

	} else {
		_, err = session.ExecCtx(ctx, query, 0, time.Now(), orderNumberId, userId)
	}
	return err

}

// SelectReceivedOrderListByShop 查看未接订单
func (o *customOrderNumberModel) SelectReceivedOrderListByShop(ctx context.Context, session sqlx.Session, shopId string, receive int, confirmedDeliver int, limit int) (*[]OrderNumber, error) {
	query := fmt.Sprintf("select %s from %s where `shop_id` = ? && order_receive = ? && order_over = 0 && confirmed_delivery = ? limit ? , 10 for update", orderNumberRows, o.table)
	var resp []OrderNumber
	if session != nil {
		err := session.QueryRowsCtx(ctx, &resp, query, shopId, receive, confirmedDeliver, limit)
		return &resp, err
	} else {
		err := o.conn.QueryRowsCtx(ctx, &resp, query, shopId, receive, confirmedDeliver, limit)
		return &resp, err
	}

}

func (o *customOrderNumberModel) SelectOrderDataListByShopAndTime(ctx context.Context, shopId string, time time.Time) (*[]OrderNumber, error) {
	query := fmt.Sprintf("select %s from %s where `shop_id` = ? and creation_time = ?", orderNumberRows, o.table)
	var resp []OrderNumber
	err := o.conn.QueryRowsCtx(ctx, &resp, query, shopId, time)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (o *customOrderNumberModel) SelectOrderNumberByOrderId(ctx context.Context, session sqlx.Session, orderId string) (*OrderNumber, error) {
	query := fmt.Sprintf("select %s from %s where `order_number` = ? and `order_over` = 0 limit 1 for update", orderNumberRows, o.table)
	var resp OrderNumber
	if session != nil {
		err := session.QueryRowCtx(ctx, &resp, query, orderId)
		if err != nil {
			return nil, err
		}
		return &resp, nil
	} else {
		err := o.conn.QueryRowCtx(ctx, &resp, query, orderId)
		if err != nil {
			return nil, err
		}
		return &resp, nil
	}
}

func (o *customOrderNumberModel) UpDateOrderReceive(ctx context.Context, session sqlx.Session, receive int64, orderNumberId string) error {

	type Req struct {
		OrderReceive int64     `db:"order_receive"` // 确认接单
		UpdataTime   time.Time `db:"updata_time"`   // 修改时间
	}
	Rows := strings.Join(builder.RawFieldNames(&Req{}), "=?,") + "=?"
	query := fmt.Sprintf("update %s set %s where `order_number` = ?", o.table, Rows)
	var err error
	if session == nil {
		_, err = o.conn.ExecCtx(ctx, query, receive, time.Now(), orderNumberId)

	} else {
		_, err = session.ExecCtx(ctx, query, receive, time.Now(), orderNumberId)
	}
	return err

}
