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

type UpDateShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpDateShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateShopLogic {
	return &UpDateShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpDateShop 店铺信息修改
func (l *UpDateShopLogic) UpDateShop(in *ams.UpDateShopReq) (*ams.UpDateShopResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		//查看店铺信息
		shop, err2 := l.svcCtx.ShopModel.TransactSelectShop(ctx, session, in.ShopId)
		if err2 != nil {
			return err2
		}
		if shop == nil {
			return errors.New("店铺信息错误")
		}

		data := model.Shop{
			ShopId:       shop.ShopId,
			ShopName:     in.ShopName,
			ShopAddress:  in.ShopAddress,
			ShopCity:     in.ShopCity,
			ShopAdmin:    shop.ShopAdmin,
			CreationTime: shop.CreationTime,
		}

		err2 = l.svcCtx.ShopModel.TransactUpDateShop(l.ctx, session, &data)
		if err2 != nil {
			return err2
		}

		return nil

	})
	if err != nil {
		return nil, err
	}
	return &ams.UpDateShopResp{
		Code:   200,
		Msg:    "修改成功",
		ShopId: in.ShopId,
	}, nil
}
