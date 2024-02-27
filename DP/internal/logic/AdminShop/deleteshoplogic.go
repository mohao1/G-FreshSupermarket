package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShopLogic {
	return &DeleteShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteShop 删除店铺
func (l *DeleteShopLogic) DeleteShop(req *types.DeleteShopReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopReq := amsclient.DeleteShopReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
	}

	//调用RPC的服务
	shopResp, err := l.svcCtx.AmsRpcClient.DeleteShop(l.ctx, &shopReq)
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
