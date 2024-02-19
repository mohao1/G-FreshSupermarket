package BmsOrder

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceivingOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceivingOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceivingOrderLogic {
	return &ReceivingOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReceivingOrder 接单
func (l *ReceivingOrderLogic) ReceivingOrder(req *types.ReceivingOrderReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	receivingOrder := bmsclient.ReceivingOrderReq{
		StaffId:     StaffId,
		OrderNumber: req.OrderNumber,
	}

	//调用RPC的服务
	receivingOrderResp, err := l.svcCtx.BmsRpcClient.ReceivingOrder(l.ctx, &receivingOrder)
	if err != nil {
		return &types.BmsDataResp{
			Code: 200,
			Msg:  "接单错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(receivingOrderResp.Code),
		Msg:  receivingOrderResp.Msg,
		Data: receivingOrderResp.OrderNumber,
	}, nil
}
