package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopLowProductSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopLowProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopLowProductSumLogic {
	return &GetShopLowProductSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopLowProductSum 进行门店对应折扣商品数量统计
func (l *GetShopLowProductSumLogic) GetShopLowProductSum(req *types.GetShopLowProductSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopLowProductSumReq := amsclient.GetShopLowProductSumReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
	}

	//调用RPC的服务
	shopLowProductSum, err := l.svcCtx.AmsRpcClient.GetShopLowProductSum(l.ctx, &shopLowProductSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopLowProductSum.Code,
		Msg:  shopLowProductSum.Msg,
		Data: shopLowProductSum.ProductSum,
	}, nil
}
