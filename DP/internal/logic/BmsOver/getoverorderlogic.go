package BmsOver

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOverOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOverOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOverOrderLogic {
	return &GetOverOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetOverOrder 获取申请取消的订单列表
func (l *GetOverOrderLogic) GetOverOrder(req *types.GetOverOrderReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	overOrderListReq := bmsclient.OverOrderListReq{
		StaffId: StaffId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	overOrderListResp, err := l.svcCtx.BmsRpcClient.GetOverOrder(l.ctx, &overOrderListReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(overOrderListResp.Code),
		Msg:  overOrderListResp.Msg,
		Data: overOrderListResp.OrderList,
	}, nil
}
