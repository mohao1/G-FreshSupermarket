package OrderServices

import (
	"DP/rpc/Sms/smsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListLogic) GetOrderList(req *types.OrderListReq) (resp *types.DataResp, err error) {
	//获取JWT中的UserId
	userId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//调用RPC的服务
	orderList, err := l.svcCtx.SmsRpcClient.GetOrderList(l.ctx, &smsclient.GetOrderListReq{
		UserId: userId,
		Limit:  req.Limit,
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
		Data: orderList,
	}, nil
}
