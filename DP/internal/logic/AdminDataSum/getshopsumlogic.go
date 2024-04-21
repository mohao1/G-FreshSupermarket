package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSumLogic {
	return &GetShopSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopSum 店铺数量
func (l *GetShopSumLogic) GetShopSum(req *types.GetShopSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopSumReq := amsclient.GetShopSumReq{
		AdminId: AdminId,
	}

	//调用RPC的服务
	shopSum, err := l.svcCtx.AmsRpcClient.GetShopSum(l.ctx, &shopSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopSum.Code,
		Msg:  shopSum.Msg,
		Data: shopSum.ShopSum,
	}, nil
}
