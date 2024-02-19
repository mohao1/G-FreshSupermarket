package StoreServices

import (
	"DP/rpc/Sms/smsclient"
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLowProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductListLogic {
	return &GetLowProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetLowProductList 获取折扣商品列表
func (l *GetLowProductListLogic) GetLowProductList(req *types.LowProductListReq) (resp *types.DataResp, err error) {
	//进行商品数据查询
	lowProductList, err := l.svcCtx.SmsRpcClient.GetLowProductList(l.ctx, &smsclient.LowProductListReq{
		ShopId:      req.ShopId,
		ProductType: req.ProductType,
		Quantity:    req.Quantity,
		Limit:       req.Limit,
	})

	//错误处理
	if err != nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "查询出错err:" + err.Error(),
			Data: nil,
		}, nil
	}

	//数据为空处理
	if lowProductList == nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "数据查询为空",
			Data: nil,
		}, nil
	}

	//数据获取成功处理
	return &types.DataResp{
		Code: int(lowProductList.Code),
		Msg:  lowProductList.Msg,
		Data: lowProductList.ProductList,
	}, nil
}
