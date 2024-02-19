package StoreServices

import (
	"DP/rpc/Sms/smsclient"
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetProductList 获取商品列表
func (l *GetProductListLogic) GetProductList(req *types.ProductListReq) (resp *types.DataResp, err error) {
	//进行商品数据查询
	productList, err := l.svcCtx.SmsRpcClient.GetProductList(l.ctx, &smsclient.ProductListReq{
		ShopId:      req.ShopId,
		ProductType: req.ProductType,
		Quantity:    req.Quantity,
		Limit:       req.Limit,
	})

	//调用报错处理情况
	if err != nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "获取失败err:" + err.Error(),
			Data: nil,
		}, nil
	}

	//数据为空处理情况
	if productList == nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "查询数据为空",
			Data: nil,
		}, nil
	}

	//返回数据
	return &types.DataResp{
		Code: int(productList.Code),
		Msg:  productList.Msg,
		Data: productList.ProductList,
	}, nil
}
