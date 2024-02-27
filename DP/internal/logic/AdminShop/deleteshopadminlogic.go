package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShopAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteShopAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShopAdminLogic {
	return &DeleteShopAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteShopAdmin 删除店铺的管理员
func (l *DeleteShopAdminLogic) DeleteShopAdmin(req *types.DeleteShopAdminReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopAdminReq := amsclient.DeleteShopAdminReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
	}

	//调用RPC的服务
	shopAdminResp, err := l.svcCtx.AmsRpcClient.DeleteShopAdmin(l.ctx, &shopAdminReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopAdminResp.Code,
		Msg:  shopAdminResp.Msg,
		Data: shopAdminResp.StaffId,
	}, nil
}
