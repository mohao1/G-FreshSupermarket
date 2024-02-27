package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopListLogic {
	return &GetShopListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopList 获取店铺列表
func (l *GetShopListLogic) GetShopList(in *ams.GetShopListReq) (*ams.GetShopListResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询店铺信息
	shopList, err := l.svcCtx.ShopModel.SelectShopList(l.ctx, in.Limit)
	if err != nil {
		return nil, err
	}

	//创建数据信息列表
	ShopList := make([]*ams.Shop, len(*shopList))
	for i, shop := range *shopList {
		ShopList[i] = &ams.Shop{
			ShopId:       shop.ShopId,
			ShopName:     shop.ShopName,
			ShopAddress:  shop.ShopAddress,
			ShopCity:     shop.ShopCity,
			CreationTime: shop.CreationTime.String(),
		}
	}

	return &ams.GetShopListResp{
		Code:     200,
		Msg:      "获取成功",
		ShopList: ShopList,
	}, nil
}
