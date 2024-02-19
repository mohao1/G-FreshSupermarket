package StoreServices

import (
	"DP/rpc/Sms/smsclient"
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopLogic {
	return &GetShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShop 根据地址获取Shop店铺列表
func (l *GetShopLogic) GetShop(req *types.GetShopReq) (resp *types.DataResp, err error) {
	//调用RPC来实现
	shop, err := l.svcCtx.SmsRpcClient.GetShop(l.ctx, &smsclient.GetShopReq{
		City: req.City,
	})
	if err != nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "err:" + err.Error(),
			Data: nil,
		}, nil
	}

	if shop.Shops == nil {
		return &types.DataResp{
			Code: int(shop.Code),
			Msg:  "选择地区没有门店",
			Data: nil,
		}, nil
	}

	return &types.DataResp{
		Code: int(shop.Code),
		Msg:  shop.Msg,
		Data: shop.Shops,
	}, nil
}
