package Product

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

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

// GetProductList 查看普通商品列表
func (l *GetProductListLogic) GetProductList(req *types.GetProductListReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	ProductList := bmsclient.GetProductListReq{
		StaffId: StaffId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	productListResp, err := l.svcCtx.BmsRpcClient.GetProductList(l.ctx, &ProductList)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(productListResp.Code),
		Msg:  productListResp.Msg,
		Data: productListResp.Product,
	}, nil
}
