package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopProductListLogic {
	return &GetShopProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopProductList 获取门店对应普通商品列表
func (l *GetShopProductListLogic) GetShopProductList(req *types.GetShopProductListReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopProductListReq := amsclient.GetShopProductListReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	shopProductList, err := l.svcCtx.AmsRpcClient.GetShopProductList(l.ctx, &shopProductListReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopProductList.Code,
		Msg:  shopProductList.Msg,
		Data: shopProductList.ProductList,
	}, nil
}
