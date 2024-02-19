package StoreServices

import (
	"DP/rpc/Sms/smsclient"
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailedProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailedProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailedProductLogic {
	return &GetDetailedProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailedProductLogic) GetDetailedProduct(req *types.DetailedProductReq) (resp *types.DataResp, err error) {

	if req.Type != "list" && req.Type != "low" {
		return &types.DataResp{
			Code: 400,
			Msg:  "查询信息类型错误",
			Data: nil,
		}, nil
	}

	//商品信息查询
	product, err := l.svcCtx.SmsRpcClient.GetDetailedProduct(l.ctx, &smsclient.DetailedProductReq{
		ProductId: req.ProductId,
		Type:      req.Type,
	})
	if err != nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "查询出现错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	if product == nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "查询数据为空",
			Data: nil,
		}, nil
	}

	return &types.DataResp{
		Code: 200,
		Msg:  "查询成功",
		Data: product,
	}, nil
}
