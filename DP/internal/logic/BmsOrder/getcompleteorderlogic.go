package BmsOrder

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompleteOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompleteOrderLogic {
	return &GetCompleteOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompleteOrderLogic) GetCompleteOrder(req *types.GetCompleteOrderReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	completeOrder := bmsclient.CompleteOrderReq{
		StaffId: StaffId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	completeOrderResp, err := l.svcCtx.BmsRpcClient.GetCompleteOrder(l.ctx, &completeOrder)
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
		Data: completeOrderResp.OrderList,
	}, nil
}
