package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopLowProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopLowProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopLowProductListLogic {
	return &GetShopLowProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopLowProductList 获取门店对应折扣商品列表
func (l *GetShopLowProductListLogic) GetShopLowProductList(req *types.GetShopLowProductListReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopLowProductListReq := amsclient.GetShopLowProductListReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	productList, err := l.svcCtx.AmsRpcClient.GetShopLowProductList(l.ctx, &shopLowProductListReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: productList.Code,
		Msg:  productList.Msg,
		Data: productList.ProductList,
	}, nil
}
