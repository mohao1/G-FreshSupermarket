package OrderServices

import (
	"DP/rpc/Sms/smsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfirmOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfirmOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmOrderLogic {
	return &ConfirmOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfirmOrderLogic) ConfirmOrder(req *types.ConfirmOrderReq) (resp *types.DataResp, err error) {
	//获取JWT中的UserId
	userId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//调用RPC的服务
	order, err := l.svcCtx.SmsRpcClient.ConfirmOrder(l.ctx, &smsclient.ConfirmOrderReq{
		UserId:      userId,
		OrderNumber: req.OrderNumber,
	})

	//错误数据处理
	if err != nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "请求错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	//正确数据返回
	return &types.DataResp{
		Code: int(order.Code),
		Msg:  order.Msg,
		Data: order.OrderNumber,
	}, nil
}
