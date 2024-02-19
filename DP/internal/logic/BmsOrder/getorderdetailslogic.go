package BmsOrder

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailsLogic {
	return &GetOrderDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetOrderDetails 查看订单详情
func (l *GetOrderDetailsLogic) GetOrderDetails(req *types.GetOrderDetailsReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	orderDetails := bmsclient.OrderDetailsReq{
		StaffId:     StaffId,
		OrderNumber: req.OrderNumber,
	}

	//调用RPC的服务
	orderDetailsResp, err := l.svcCtx.BmsRpcClient.GetOrderDetails(l.ctx, &orderDetails)
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
		Data: orderDetailsResp,
	}, nil
}
