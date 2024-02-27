package model

import (
	"context"
	"database/sql"
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
		TransactInsertStaff(ctx context.Context, session sqlx.Session, data *Staff) (sql.Result, error)
		TransactDeleteStaffPT(ctx context.Context, session sqlx.Session, shopId string, positionId string) error
		TransactSelectAllStaffList(ctx context.Context, session sqlx.Session, limit int64) (*[]Staff, error)
		SelectStaffCount(ctx context.Context) (*[]StaffCount, error)
	}

	customStaffModel struct {
		*defaultStaffModel
	}

	StaffCount struct {
		ShopId      string `db:"shop_id"`      // 店铺id
		CountNumber int64  `db:"count_number"` // 统计人数数值
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

func (c *customStaffModel) TransactInsertStaff(ctx context.Context, session sqlx.Session, data *Staff) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", c.table, staffRowsExpectAutoSet)
	if session != nil {
		ret, err := session.ExecCtx(ctx, query, data.StaffId, data.StaffName, data.PositionId, data.Password, data.ShopId, data.CreationTime)
		return ret, err
	} else {
		ret, err := c.conn.ExecCtx(ctx, query, data.StaffId, data.StaffName, data.PositionId, data.Password, data.ShopId, data.CreationTime)
		return ret, err
	}
}

// TransactDeleteStaffPT 删除所有的普通的员工
func (c *customStaffModel) TransactDeleteStaffPT(ctx context.Context, session sqlx.Session, shopId string, positionId string) error {
	query := fmt.Sprintf("delete from %s where `shop_id` = ? and `position_id` = ?", c.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, shopId, positionId)
		return err
	} else {
		_, err := c.conn.ExecCtx(ctx, query, shopId, positionId)
		return err
	}
}

// TransactSelectAllStaffList 查询全部员工
func (c *customStaffModel) TransactSelectAllStaffList(ctx context.Context, session sqlx.Session, limit int64) (*[]Staff, error) {
	query := fmt.Sprintf("select %s from %s limit ? , 10 for update", staffRows, c.table)
	var resp []Staff
	if session != nil {
		err := session.QueryRowsCtx(ctx, &resp, query, limit)
		if err != nil {
			return nil, err
		}
	} else {
		err := c.conn.QueryRowsCtx(ctx, &resp, query, limit)
		if err != nil {
			return nil, err
		}
	}

	return &resp, nil

}

func (c *customStaffModel) SelectStaffCount(ctx context.Context) (*[]StaffCount, error) {
	query := fmt.Sprintf("SELECT `shop_id` , COUNT(*) AS `count_number` FROM %s WHERE shop_id != \"\"   GROUP BY `shop_id`;", c.table)
	var resp []StaffCount
	err := c.conn.QueryRowsCtx(ctx, &resp, query)
	return &resp, err
}
