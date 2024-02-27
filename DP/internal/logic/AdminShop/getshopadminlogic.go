package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopAdminLogic {
	return &GetShopAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopAdmin 查看店铺的管理员
func (l *GetShopAdminLogic) GetShopAdmin(req *types.GetShopAdminReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopAdminReq := amsclient.GetShopAdminReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
	}

	//调用RPC的服务
	shopAdminResp, err := l.svcCtx.AmsRpcClient.GetShopAdmin(l.ctx, &shopAdminReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: 200,
		Msg:  "获取成功",
		Data: shopAdminResp,
	}, nil
}
