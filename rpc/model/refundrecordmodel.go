package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ RefundRecordModel = (*customRefundRecordModel)(nil)

type (
	// RefundRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRefundRecordModel.
	RefundRecordModel interface {
		refundRecordModel
		TransactInsert(ctx context.Context, session sqlx.Session, data *RefundRecord) (sql.Result, error)
		TransactUpDataRefundOver(ctx context.Context, session sqlx.Session, orderId string, userId string) error
		SelectRefundRecordListByShop(ctx context.Context, shopId string, limit int, refundOver int, confirm int) (*[]RefundRecord, error)
		TransactUpDateOverIsOk(ctx context.Context, session sqlx.Session, shopId string, orderId string) error
		TransactUpDateOverIsNo(ctx context.Context, session sqlx.Session, shopId string, orderId string) error
	}

	customRefundRecordModel struct {
		*defaultRefundRecordModel
	}
)

// NewRefundRecordModel returns a model for the database table.
func NewRefundRecordModel(conn sqlx.SqlConn) RefundRecordModel {
	return &customRefundRecordModel{
		defaultRefundRecordModel: newRefundRecordModel(conn),
	}
}

func (r *customRefundRecordModel) TransactInsert(ctx context.Context, session sqlx.Session, data *RefundRecord) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", r.table, refundRecordRowsExpectAutoSet)
	if session == nil {
		ret, err := r.conn.ExecCtx(ctx, query, data.RefundId, data.ShopId, data.OrderId, data.UserId, data.RefundOver, data.Confirm, data.DeleteKey)
		return ret, err
	} else {
		ret, err := session.ExecCtx(ctx, query, data.RefundId, data.ShopId, data.OrderId, data.UserId, data.RefundOver, data.Confirm, data.DeleteKey)
		return ret, err
	}
}

func (r *customRefundRecordModel) TransactUpDataRefundOver(ctx context.Context, session sqlx.Session, orderId string, userId string) error {

	type Req struct {
		RefundOver int64 `db:"refund_over"` // 是否取消申请
	}
	Rows := strings.Join(builder.RawFieldNames(&Req{}), "=?,") + "=?"
	query := fmt.Sprintf("update %s set %s where `order_id` = ? AND `user_id` = ? ", r.table, Rows)

	var err error
	if session == nil {
		_, err = r.conn.ExecCtx(ctx, query, 1, orderId, userId)

	} else {
		_, err = session.ExecCtx(ctx, query, 1, orderId, userId)
	}

	return err
}

func (r *customRefundRecordModel) SelectRefundRecordListByShop(ctx context.Context, shopId string, limit int, refundOver int, confirm int) (*[]RefundRecord, error) {
	query := fmt.Sprintf("select %s from %s where `shop_id` = ? AND `refund_over` = ? AND `confirm` = ? limit ? , 10", refundRecordRows, r.table)
	var resp []RefundRecord
	err := r.conn.QueryRowsCtx(ctx, &resp, query, shopId, refundOver, confirm, limit)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *customRefundRecordModel) TransactUpDateOverIsOk(ctx context.Context, session sqlx.Session, shopId string, orderId string) error {
	type Req struct {
		Confirm int64 `db:"confirm"` // 是否同意
	}
	Rows := strings.Join(builder.RawFieldNames(&Req{}), "=?,") + "=?"
	query := fmt.Sprintf("update %s set %s where `order_id` = ? AND `shop_id` = ? ", r.table, Rows)

	var err error
	if session == nil {
		_, err = r.conn.ExecCtx(ctx, query, 1, orderId, shopId)

	} else {
		_, err = session.ExecCtx(ctx, query, 1, orderId, shopId)
	}

	return err
}

func (r *customRefundRecordModel) TransactUpDateOverIsNo(ctx context.Context, session sqlx.Session, shopId string, orderId string) error {
	type Req struct {
		Confirm int64 `db:"confirm"` // 是否同意
	}
	Rows := strings.Join(builder.RawFieldNames(&Req{}), "=?,") + "=?"
	query := fmt.Sprintf("update %s set %s where `order_id` = ? AND `shop_id` = ? ", r.table, Rows)

	var err error
	if session == nil {
		_, err = r.conn.ExecCtx(ctx, query, 2, orderId, shopId)

	} else {
		_, err = session.ExecCtx(ctx, query, 2, orderId, shopId)
	}

	return err
}
