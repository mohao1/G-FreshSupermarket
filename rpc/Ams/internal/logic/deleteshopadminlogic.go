package logic

import (
	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"
	"DP/rpc/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShopAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteShopAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShopAdminLogic {
	return &DeleteShopAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteShopAdmin 删除店铺的管理员
func (l *DeleteShopAdminLogic) DeleteShopAdmin(in *ams.DeleteShopAdminReq) (*ams.DeleteShopAdminResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	staffId := ""

	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		//查询是否商店存在
		shop, err2 := l.svcCtx.ShopModel.TransactSelectShop(ctx, session, in.ShopId)
		if err2 != nil {
			return err2
		}
		if shop == nil {
			return errors.New("商店信息错误")
		}

		//判断是否没有商店管理人员
		if shop.ShopAdmin == "" {
			return errors.New("商店没有存在管理人员")
		}

		//记录角色Id
		staffId = shop.ShopAdmin

		//查询管理人员信息
		staff, err2 := l.svcCtx.StaffModel.TransactSelectStaff(l.ctx, session, shop.ShopAdmin)
		if err2 != nil {
			return err2
		}
		if staff == nil {
			return errors.New("用户的账号不存在")
		}

		//判断管理人员是否管理对应商店
		if staff.ShopId != in.ShopId {
			return errors.New("管理信息错误")
		}

		data := model.Staff{
			StaffId:      staff.StaffId,
			StaffName:    staff.StaffName,
			PositionId:   staff.PositionId,
			Password:     staff.Password,
			ShopId:       "",
			CreationTime: staff.CreationTime,
		}

		//添加管理账户捆绑信息
		err2 = l.svcCtx.StaffModel.TransactUpDateStaff(ctx, session, &data)
		if err2 != nil {
			return err2
		}

		//设置Shop的管理人员信息
		shopData := model.Shop{
			ShopId:       in.ShopId,
			ShopName:     shop.ShopName,
			ShopAddress:  shop.ShopAddress,
			ShopCity:     shop.ShopCity,
			ShopAdmin:    "",
			CreationTime: shop.CreationTime,
		}
		//解绑商店管理人员信息
		err2 = l.svcCtx.ShopModel.TransactUpDateShop(ctx, session, &shopData)
		if err2 != nil {
			return err2
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &ams.DeleteShopAdminResp{
		Code:    200,
		Msg:     "删除成功",
		StaffId: staffId,
	}, nil
}
