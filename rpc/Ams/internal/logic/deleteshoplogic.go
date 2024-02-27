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

type DeleteShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShopLogic {
	return &DeleteShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteShop 删除店铺
func (l *DeleteShopLogic) DeleteShop(in *ams.DeleteShopReq) (*ams.DeleteShopResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//查询店铺
		shop, err2 := l.svcCtx.ShopModel.TransactSelectShop(ctx, session, in.ShopId)
		if err2 != nil {
			return err2
		}
		//查看是否存在
		if shop == nil {
			return errors.New("店铺信息错误")
		}

		if shop.ShopAdmin != "" {
			//判断是否相同
			staffAdmin, err3 := l.svcCtx.StaffModel.TransactSelectStaff(ctx, session, shop.ShopAdmin)
			if err3 != nil {
				return err3
			}

			if staffAdmin.ShopId != shop.ShopId {
				return errors.New("商店和管理的信息错误")
			}

			//解除管理人员绑定
			staffData := model.Staff{
				StaffId:      staffAdmin.StaffId,
				StaffName:    staffAdmin.StaffName,
				PositionId:   staffAdmin.PositionId,
				Password:     staffAdmin.Password,
				ShopId:       "",
				CreationTime: staffAdmin.CreationTime,
			}
			err3 = l.svcCtx.StaffModel.TransactUpDateStaff(ctx, session, &staffData)
			if err3 != nil {
				return err3
			}

		}

		//删除所有的普通员工
		err2 = l.svcCtx.StaffModel.TransactDeleteStaffPT(ctx, session, shop.ShopId, "2")
		if err2 != nil {
			return err2
		}

		//删除对应店铺
		err2 = l.svcCtx.ShopModel.TransactDeleteShop(ctx, session, in.ShopId)
		if err2 != nil {
			return err2
		}

		return nil

	})
	if err != nil {
		return nil, err
	}

	return &ams.DeleteShopResp{
		Code:   200,
		Msg:    "删除成功",
		ShopId: in.ShopId,
	}, nil
}
