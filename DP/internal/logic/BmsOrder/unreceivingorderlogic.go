package BmsOrder

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnReceivingOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnReceivingOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnReceivingOrderLogic {
	return &UnReceivingOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UnReceivingOrder 取消接单
func (l *UnReceivingOrderLogic) UnReceivingOrder(req *types.UnReceivingOrderReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	unReceivingOrder := bmsclient.UnReceivingOrderReq{
		StaffId:     StaffId,
		OrderNumber: req.OrderNumber,
	}

	//调用RPC的服务
	unReceivingOrderResp, err := l.svcCtx.BmsRpcClient.UnReceivingOrder(l.ctx, &unReceivingOrder)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "取消接单错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(unReceivingOrderResp.Code),
		Msg:  unReceivingOrderResp.Msg,
		Data: unReceivingOrderResp.OrderNumber,
	}, nil
}
