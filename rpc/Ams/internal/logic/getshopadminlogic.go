package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopAdminLogic {
	return &GetShopAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopAdmin 查看店铺的管理员
func (l *GetShopAdminLogic) GetShopAdmin(in *ams.GetShopAdminReq) (*ams.GetShopAdminResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	shop, err := l.svcCtx.ShopModel.TransactSelectShop(l.ctx, nil, in.ShopId)
	if err != nil {
		return nil, err
	}

	//判断商店信息是否存在
	if shop == nil {
		return nil, errors.New("商店信息错误")
	}

	//判断商店是否存在管理人员
	if shop.ShopAdmin == "" {
		return nil, errors.New("店铺没有管理人员")
	}

	//查询管理人员信息
	staff, err := l.svcCtx.StaffModel.TransactSelectStaff(l.ctx, nil, shop.ShopAdmin)
	if err != nil {
		return nil, err
	}

	//判断管理人员信息是否正确
	if shop == nil {
		return nil, errors.New("管理人员信息错误")
	}

	//查看职业是否存在
	position, err := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
	if err != nil {
		return nil, err
	}

	return &ams.GetShopAdminResp{
		StaffId:      staff.StaffId,
		StaffName:    staff.StaffName,
		PositionName: position.PositionName,
		ShopId:       staff.ShopId,
		ShopName:     shop.ShopName,
		CreationTime: staff.CreationTime.String(),
	}, nil
}
