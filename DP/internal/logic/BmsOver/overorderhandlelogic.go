package BmsOver

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OverOrderHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOverOrderHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OverOrderHandleLogic {
	return &OverOrderHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// OverOrderHandle 取消订单申请处理
func (l *OverOrderHandleLogic) OverOrderHandle(req *types.OverOrderHandleReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	overOrderHandleReq := bmsclient.OverOrderHandleReq{
		StaffId:     StaffId,
		OrderNumber: req.OrderNumber,
		Type:        req.TypeNumber,
	}

	//调用RPC的服务
	overOrderHandleResp, err := l.svcCtx.BmsRpcClient.OverOrderHandle(l.ctx, &overOrderHandleReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(overOrderHandleResp.Code),
		Msg:  overOrderHandleResp.Msg,
		Data: overOrderHandleResp.ProductId,
	}, nil
}
