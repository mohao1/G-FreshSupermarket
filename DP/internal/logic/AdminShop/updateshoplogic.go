package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateShopLogic {
	return &UpDateShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpDateShop 店铺信息修改
func (l *UpDateShopLogic) UpDateShop(req *types.UpDateShopReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopReq := amsclient.UpDateShopReq{
		AdminId:     AdminId,
		ShopId:      req.ShopId,
		ShopName:    req.ShopName,
		ShopAddress: req.ShopAddress,
		ShopCity:    req.ShopCity,
	}

	//调用RPC的服务
	shopResp, err := l.svcCtx.AmsRpcClient.UpDateShop(l.ctx, &shopReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopResp.Code,
		Msg:  shopResp.Msg,
		Data: shopResp.ShopId,
	}, nil
}
