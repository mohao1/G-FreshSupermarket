package BmsOrder

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReceivedOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReceivedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReceivedOrderLogic {
	return &GetReceivedOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetReceivedOrder 查看已接订单
func (l *GetReceivedOrderLogic) GetReceivedOrder(req *types.GetReceivedOrderReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	receivedOrder := bmsclient.ReceivedOrderReq{
		StaffId: StaffId,
		Limit:   req.Limit,
	}
	//调用RPC的服务
	receivedOrderResp, err := l.svcCtx.BmsRpcClient.GetReceivedOrder(l.ctx, &receivedOrder)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "获取信息错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: 200,
		Msg:  "获取成功",
		Data: receivedOrderResp.OrderList,
	}, nil
}
