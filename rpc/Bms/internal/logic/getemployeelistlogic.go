package logic

import (
	"context"
	"errors"
	"sync"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmployeeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewGetEmployeeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmployeeListLogic {
	return &GetEmployeeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetEmployeeList 获取员工列表
func (l *GetEmployeeListLogic) GetEmployeeList(in *bms.GetEmployeeListReq) (*bms.GetEmployeeListResp, error) {

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

	//获取数据信息列表
	staffList, err := l.svcCtx.StaffModel.TransactSelectStaffList(l.ctx, nil, staff.ShopId)
	if err != nil {

		return nil, err
	}

	//定义数据信息
	newStaffList := make([]*bms.Staff, len(*staffList))

	//设置数据信息
	for i, staffL := range *staffList {

		//查询权限信息
		positionL, err2 := l.svcCtx.PositionModel.FindOne(l.ctx, staffL.PositionId)
		if err2 != nil {
			return nil, err2
		}

		shop, err2 := l.svcCtx.ShopModel.FindOne(l.ctx, staffL.ShopId)
		if err2 != nil {
			return nil, err2
		}
		if err2 != nil {
			return nil, err2
		}

		//设置数据信息
		newStaffList[i] = &bms.Staff{
			StaffId:      staffL.StaffId,
			StaffName:    staffL.StaffName,
			PositionId:   positionL.PositionName,
			PassWord:     staffL.Password,
			ShopName:     shop.ShopName,
			CreationTime: staffL.CreationTime.String(),
		}
	}

	return &bms.GetEmployeeListResp{
		Code:      200,
		Msg:       "获取成功",
		StaffList: newStaffList,
	}, nil
}
