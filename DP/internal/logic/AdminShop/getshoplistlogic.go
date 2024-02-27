package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopListLogic {
	return &GetShopListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopList 获取店铺列表
func (l *GetShopListLogic) GetShopList(req *types.GetShopListReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopListReq := amsclient.GetShopListReq{
		AdminId: AdminId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	shopListResp, err := l.svcCtx.AmsRpcClient.GetShopList(l.ctx, &shopListReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopListResp.Code,
		Msg:  shopListResp.Msg,
		Data: shopListResp.ShopList,
	}, nil
}
