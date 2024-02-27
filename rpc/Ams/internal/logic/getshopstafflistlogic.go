package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopStaffListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopStaffListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopStaffListLogic {
	return &GetShopStaffListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopStaffList 店铺查询人员列表
func (l *GetShopStaffListLogic) GetShopStaffList(in *ams.GetShopStaffListReq) (*ams.GetShopStaffListResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	staffs, err := l.svcCtx.StaffModel.TransactSelectStaffList(l.ctx, nil, in.ShopId)
	if err != nil {
		return nil, err
	}

	//准备数据
	staffList := make([]*ams.Staff, len(*staffs))

	for i, staff := range *staffs {
		shop, err2 := l.svcCtx.ShopModel.FindOne(l.ctx, staff.ShopId)
		if err2 != nil {
			return nil, err2
		}

		position, err2 := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
		if err2 != nil {
			return nil, err2
		}

		staffList[i] = &ams.Staff{
			StaffId:      staff.StaffId,
			StaffName:    staff.StaffName,
			ShopId:       staff.ShopId,
			ShopName:     shop.ShopName,
			PositionName: staff.PositionId,
			PositionId:   position.PositionName,
			CreationTime: staff.CreationTime.String(),
		}
	}

	return &ams.GetShopStaffListResp{
		Code:      200,
		Msg:       "获取成功",
		StaffList: staffList,
	}, nil
}
