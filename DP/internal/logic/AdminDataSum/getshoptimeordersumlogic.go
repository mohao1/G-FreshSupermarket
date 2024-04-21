package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopTimeOrderSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopTimeOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopTimeOrderSumLogic {
	return &GetShopTimeOrderSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopTimeOrderSum 各个店铺根据时间段的订单数量
func (l *GetShopTimeOrderSumLogic) GetShopTimeOrderSum(req *types.GetShopTimeOrderSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopTimeOrderSumReq := amsclient.GetShopTimeOrderSumReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
		TopTime: req.TopTime,
		EndTime: req.EndTime,
	}

	//调用RPC的服务
	shopTimeOrderSum, err := l.svcCtx.AmsRpcClient.GetShopTimeOrderSum(l.ctx, &shopTimeOrderSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopTimeOrderSum.Code,
		Msg:  shopTimeOrderSum.Msg,
		Data: shopTimeOrderSum.ShopTimeOrderList,
	}, nil
}
