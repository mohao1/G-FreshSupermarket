package logic

import (
	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllShopStaffListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllShopStaffListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllShopStaffListLogic {
	return &GetAllShopStaffListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetAllShopStaffList 查看全部店铺员工列表
func (l *GetAllShopStaffListLogic) GetAllShopStaffList(in *ams.GetAllShopStaffListReq) (*ams.GetAllShopStaffListResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询员工信息数据
	staffs, err := l.svcCtx.StaffModel.TransactSelectAllStaffList(l.ctx, nil, in.Limit)
	if err != nil {
		return nil, err
	}

	//准备数据
	staffList := make([]*ams.Staff, len(*staffs))

	for i, staff := range *staffs {
		var shopName string
		if staff.ShopId != "" {
			shop, err2 := l.svcCtx.ShopModel.FindOne(l.ctx, staff.ShopId)
			if err2 != nil {
				return nil, err2
			}
			shopName = shop.ShopName
		}

		position, err2 := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
		if err2 != nil {
			return nil, err2
		}

		staffList[i] = &ams.Staff{
			StaffId:      staff.StaffId,
			StaffName:    staff.StaffName,
			ShopId:       staff.ShopId,
			ShopName:     shopName,
			PositionName: position.PositionName,
			PositionId:   staff.PositionId,
			CreationTime: staff.CreationTime.String(),
		}
	}

	return &ams.GetAllShopStaffListResp{
		Code:      200,
		Msg:       "获取成功",
		StaffList: staffList,
	}, nil
}
