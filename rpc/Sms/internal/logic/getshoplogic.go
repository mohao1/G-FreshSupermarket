package logic

import (
	"context"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopLogic {
	return &GetShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShop 获取店铺
func (l *GetShopLogic) GetShop(in *sms.GetShopReq) (*sms.GetShopResp, error) {
	city, err := l.svcCtx.ShopModel.SelectShopByCity(l.ctx, in.City)
	if err != nil {
		return nil, err
	}
	var shops = make([]*sms.Shop, len(*city))
	for i, shop := range *city {
		shops[i] = &sms.Shop{
			ShopId:   shop.ShopId,
			ShopName: shop.ShopName,
			Address:  shop.ShopAddress,
			City:     shop.ShopCity,
		}
	}
	return &sms.GetShopResp{
		Code:  200,
		Shops: shops,
		Msg:   "获取成功",
	}, nil
}
