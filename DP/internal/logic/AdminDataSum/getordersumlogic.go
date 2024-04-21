package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderSumLogic {
	return &GetOrderSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetOrderSum 今日消费用户数量
func (l *GetOrderSumLogic) GetOrderSum(req *types.GetOrderSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	orderSumReq := amsclient.GetOrderSumReq{
		AdminId: AdminId,
	}

	//调用RPC的服务
	orderSum, err := l.svcCtx.AmsRpcClient.GetOrderSum(l.ctx, &orderSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: 200,
		Msg:  "获取成功",
		Data: orderSum.OrderSum,
	}, nil
}
