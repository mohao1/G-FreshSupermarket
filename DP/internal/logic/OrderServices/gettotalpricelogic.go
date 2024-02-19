package OrderServices

import (
	"DP/DP/internal/svc"
	"DP/DP/internal/types"
	"DP/rpc/Sms/smsclient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type GetTotalPriceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTotalPriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTotalPriceLogic {
	return &GetTotalPriceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTotalPriceLogic) GetTotalPrice(req *types.TotalPriceReq) (resp *types.DataResp, err error) {

	//调用RPC的服务
	price, err := l.svcCtx.SmsRpcClient.GetTotalPrice(l.ctx, &smsclient.GetTotalPriceReq{
		ProductId:       req.ProductId,
		ProductQuantity: req.ProductQuantity,
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
	coed, _ := strconv.Atoi(price.Code)
	return &types.DataResp{
		Code: coed,
		Msg:  price.Msg,
		Data: price.TotalPrice,
	}, nil
}
