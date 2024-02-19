package OrderServices

import (
	"DP/rpc/Sms/smsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailedOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailedOrderLogic {
	return &GetDetailedOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailedOrderLogic) GetDetailedOrder(req *types.GetDetailedOrderReq) (resp *types.DataResp, err error) {
	//获取JWT中的UserId
	userId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//调用RPC的服务
	order, err := l.svcCtx.SmsRpcClient.GetDetailedOrder(l.ctx, &smsclient.DetailedOrderReq{
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
		Code: 200,
		Msg:  "数据请求成功",
		Data: order,
	}, nil
}
