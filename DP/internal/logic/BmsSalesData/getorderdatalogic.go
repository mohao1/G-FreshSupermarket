package BmsSalesData

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDataLogic {
	return &GetOrderDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetOrderData 查看订单数据
func (l *GetOrderDataLogic) GetOrderData(req *types.GetOrderDataReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	orderDataReq := bmsclient.OrderDataReq{
		StaffId: StaffId,
		GetTime: req.GetTime,
	}

	//调用RPC的服务
	orderDataResp, err := l.svcCtx.BmsRpcClient.GetOrderData(l.ctx, &orderDataReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(orderDataResp.Code),
		Msg:  orderDataResp.Msg,
		Data: orderDataResp.OrderData,
	}, nil
}
