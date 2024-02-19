package BmsOrder

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnreceivedOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUnreceivedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnreceivedOrderLogic {
	return &GetUnreceivedOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUnreceivedOrder 查看未接订单
func (l *GetUnreceivedOrderLogic) GetUnreceivedOrder(req *types.GetUnreceivedOrderReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	GetUnreceivedOrder := bmsclient.UnreceivedOrderReq{
		StaffId: StaffId,
		Limit:   req.Limit,
	}
	//调用RPC的服务
	unreceivedOrder, err := l.svcCtx.BmsRpcClient.GetUnreceivedOrder(l.ctx, &GetUnreceivedOrder)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "查询错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: 200,
		Msg:  "获取成功",
		Data: unreceivedOrder,
	}, nil
}
