package logic

import (
	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"
	"DP/rpc/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type PostShopAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostShopAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostShopAdminLogic {
	return &PostShopAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PostShopAdmin 设置店铺的管理员
func (l *PostShopAdminLogic) PostShopAdmin(in *ams.PostShopAdminReq) (*ams.PostShopAdminResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return &ams.PostShopAdminResp{
			Code:    400,
			Msg:     "权限不足",
			StaffId: "",
		}, nil
	}

	//检测用户是否存在
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if staff == nil {
		return &ams.PostShopAdminResp{
			Code:    400,
			Msg:     "用户没有在",
			StaffId: "",
		}, nil
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		//查看商店是否存在
		shop, err2 := l.svcCtx.ShopModel.TransactSelectShop(ctx, session, in.ShopId)
		if err2 != nil {
			return err2
		}
		if shop == nil {
			return errors.New("商店信息错误")
		}

		//判断是否已经存在商店管理人员
		if shop.ShopAdmin != "" {
			return errors.New("商店已经存在管理人员")
		}

		//查看账号是否已经存在管理店铺
		Staff, err2 := l.svcCtx.StaffModel.TransactSelectStaff(ctx, session, in.StaffId)
		if err2 != nil {
			return err2
		}
		if Staff == nil {
			return errors.New("用户的账号不存在")
		}
		if Staff.ShopId != "" {
			return errors.New("用户的账号已经存在管理商店")
		}

		staffData := model.Staff{
			StaffId:      staff.StaffId,
			StaffName:    staff.StaffName,
			PositionId:   staff.PositionId,
			Password:     staff.Password,
			ShopId:       shop.ShopId,
			CreationTime: staff.CreationTime,
		}

		//设置店铺的管理员信息
		err2 = l.svcCtx.StaffModel.TransactUpDateStaff(ctx, session, &staffData)
		if err2 != nil {
			return err2
		}

		//设置Shop的管理人员信息
		shopData := model.Shop{
			ShopId:       in.ShopId,
			ShopName:     shop.ShopName,
			ShopAddress:  shop.ShopAddress,
			ShopCity:     shop.ShopCity,
			ShopAdmin:    in.StaffId,
			CreationTime: shop.CreationTime,
		}
		err2 = l.svcCtx.ShopModel.TransactUpDateShop(ctx, session, &shopData)
		if err2 != nil {
			return err2
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &ams.PostShopAdminResp{
		Code:    200,
		Msg:     "设置成功",
		StaffId: in.StaffId,
	}, nil
}
