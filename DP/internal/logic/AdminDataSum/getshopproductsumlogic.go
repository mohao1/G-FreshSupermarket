package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopProductSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopProductSumLogic {
	return &GetShopProductSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopProductSum 进行门店对应普通商品数量统计
func (l *GetShopProductSumLogic) GetShopProductSum(req *types.GetShopProductSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopProductSumReq := amsclient.GetShopProductSumReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
	}

	//调用RPC的服务
	shopProductSum, err := l.svcCtx.AmsRpcClient.GetShopProductSum(l.ctx, &shopProductSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopProductSum.Code,
		Msg:  shopProductSum.Msg,
		Data: shopProductSum.ProductSum,
	}, nil
}
