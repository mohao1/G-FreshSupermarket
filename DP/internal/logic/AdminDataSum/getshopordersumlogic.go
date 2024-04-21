package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopOrderSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopOrderSumLogic {
	return &GetShopOrderSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopOrderSum 各个店铺总的订单数量
func (l *GetShopOrderSumLogic) GetShopOrderSum(req *types.GetShopOrderSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopOrderSumReq := amsclient.GetShopOrderSumReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
	}

	//调用RPC的服务
	shopOrderSum, err := l.svcCtx.AmsRpcClient.GetShopOrderSum(l.ctx, &shopOrderSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopOrderSum.Code,
		Msg:  shopOrderSum.Msg,
		Data: shopOrderSum.ShopOrderSumList,
	}, nil
}
