package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NoticeModel = (*customNoticeModel)(nil)

type (
	// NoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNoticeModel.
	NoticeModel interface {
		noticeModel
		SelectNoticeByShopId(ctx context.Context, shopId string) (*[]Notice, error)
	}

	customNoticeModel struct {
		*defaultNoticeModel
	}
)

// NewNoticeModel returns a model for the database table.
func NewNoticeModel(conn sqlx.SqlConn) NoticeModel {
	return &customNoticeModel{
		defaultNoticeModel: newNoticeModel(conn),
	}
}

func (c *customNoticeModel) SelectNoticeByShopId(ctx context.Context, shopId string) (*[]Notice, error) {
	query := fmt.Sprintf("select %s from %s where `shop_id` = ?", noticeRows, c.table)
	var resp []Notice
	err := c.conn.QueryRowsCtx(ctx, &resp, query, shopId)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
