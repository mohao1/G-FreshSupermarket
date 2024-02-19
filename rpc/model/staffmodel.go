package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StaffModel = (*customStaffModel)(nil)

type (
	// StaffModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStaffModel.
	StaffModel interface {
		staffModel
		TransactSelectStaff(ctx context.Context, session sqlx.Session, staffId string) (*Staff, error)
		TransactUpDateStaff(ctx context.Context, session sqlx.Session, data *Staff) error
		TransactDeleteStaff(ctx context.Context, session sqlx.Session, staffId string) error
		TransactSelectStaffList(ctx context.Context, session sqlx.Session, shopId string) (*[]Staff, error)
	}

	customStaffModel struct {
		*defaultStaffModel
	}
)

// NewStaffModel returns a model for the database table.
func NewStaffModel(conn sqlx.SqlConn) StaffModel {
	return &customStaffModel{
		defaultStaffModel: newStaffModel(conn),
	}
}

func (c *customStaffModel) TransactSelectStaff(ctx context.Context, session sqlx.Session, staffId string) (*Staff, error) {
	query := fmt.Sprintf("select %s from %s where `staff_id` = ? limit 1 for update", staffRows, c.table)
	var resp Staff
	if session != nil {
		err := session.QueryRowCtx(ctx, &resp, query, staffId)
		if err != nil {
			return nil, err
		}
	} else {
		err := c.conn.QueryRowCtx(ctx, &resp, query, staffId)
		if err != nil {
			return nil, err
		}
	}

	return &resp, nil
}

func (c *customStaffModel) TransactUpDateStaff(ctx context.Context, session sqlx.Session, data *Staff) error {
	query := fmt.Sprintf("update %s set %s where `staff_id` = ?", c.table, staffRowsWithPlaceHolder)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, data.StaffName, data.PositionId, data.Password, data.ShopId, data.CreationTime, data.StaffId)
		return err
	} else {
		_, err := c.conn.ExecCtx(ctx, query, data.StaffName, data.PositionId, data.Password, data.ShopId, data.CreationTime, data.StaffId)
		return err
	}
}

func (c *customStaffModel) TransactDeleteStaff(ctx context.Context, session sqlx.Session, staffId string) error {
	query := fmt.Sprintf("delete from %s where `staff_id` = ?", c.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, staffId)
		return err
	} else {
		_, err := c.conn.ExecCtx(ctx, query, staffId)
		return err
	}
}

func (c *customStaffModel) TransactSelectStaffList(ctx context.Context, session sqlx.Session, shopId string) (*[]Staff, error) {
	query := fmt.Sprintf("select %s from %s where `shop_id` = ? for update", staffRows, c.table)
	var resp []Staff
	if session != nil {
		err := session.QueryRowsCtx(ctx, &resp, query, shopId)
		if err != nil {
			return nil, err
		}
	} else {
		err := c.conn.QueryRowsCtx(ctx, &resp, query, shopId)
		if err != nil {
			return nil, err
		}
	}

	return &resp, nil
}
