package logic

import (
	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"
	"DP/rpc/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateEmployeeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewUpDateEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateEmployeeLogic {
	return &UpDateEmployeeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpDateEmployee 修改员工信息
func (l *UpDateEmployeeLogic) UpDateEmployee(in *bms.UpDateEmployeeReq) (*bms.UpDateEmployeeResp, error) {

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

	if in.PositionId == position.Id {
		return nil, errors.New("一个店铺不能重复设置经理")
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

		data := model.Staff{
			StaffId:      in.SetStaffId,
			StaffName:    in.StaffName,
			PositionId:   in.PositionId,
			Password:     in.PassWord,
			ShopId:       selectStaff.ShopId,
			CreationTime: selectStaff.CreationTime,
		}

		err2 = l.svcCtx.StaffModel.TransactUpDateStaff(l.ctx, session, &data)
		if err2 != nil {
			return err2
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &bms.UpDateEmployeeResp{
		Code:    200,
		Msg:     "修改成功",
		StaffId: in.SetStaffId,
	}, nil
}
