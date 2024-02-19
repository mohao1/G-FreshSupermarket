package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"sync"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmployeeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewDeleteEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmployeeLogic {
	return &DeleteEmployeeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteEmployee 删除员工
func (l *DeleteEmployeeLogic) DeleteEmployee(in *bms.DeleteEmployeeReq) (*bms.DeleteEmployeeResp, error) {

	//查询店铺信息和个人的信息
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//进行权限判断
	position, err := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
	if err != nil {
		return nil, err
	}
	if position.PositionName != "经理" {
		return nil, errors.New("权限不足")
	}

	//上锁
	l.rWMutex.Lock()
	defer l.rWMutex.Unlock()

	//开启事务
	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		selectStaff, err2 := l.svcCtx.StaffModel.TransactSelectStaff(ctx, session, in.SetStaffId)
		if err2 != nil {
			return err2
		}

		if selectStaff.ShopId != staff.ShopId {
			return errors.New("店铺信息错误没有权限删除员工")
		}

		err2 = l.svcCtx.StaffModel.TransactDeleteStaff(l.ctx, session, in.SetStaffId)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &bms.DeleteEmployeeResp{
		Code:    200,
		Msg:     "删除成功",
		StaffId: in.SetStaffId,
	}, nil
}
