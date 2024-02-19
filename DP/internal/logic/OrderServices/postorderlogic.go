package OrderServices

import (
	"DP/DP/internal/svc"
	"DP/DP/internal/types"
	"DP/rpc/Sms/smsclient"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostOrderLogic {
	return &PostOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostOrderLogic) PostOrder(req *types.PostOrderReq) (resp *types.DataResp, err error) {
	//获取JWT中的UserId
	userId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//调用RPC的服务
	order, err := l.svcCtx.SmsRpcClient.PostOrder(l.ctx, &smsclient.PostOrderReq{
		UserId:          userId,
		ShopId:          req.ShopId,
		ProductId:       req.ProductId,
		ProductQuantity: req.ProductQuantity,
		DeliveryTime:    req.DeliveryTime,
		Notes:           req.Notes,
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
