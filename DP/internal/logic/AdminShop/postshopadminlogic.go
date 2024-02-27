package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostShopAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostShopAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostShopAdminLogic {
	return &PostShopAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PostShopAdmin 设置店铺的管理员
func (l *PostShopAdminLogic) PostShopAdmin(req *types.PostShopAdminReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	adminReq := amsclient.PostShopAdminReq{
		AdminId: AdminId,
		StaffId: req.StaffId,
		ShopId:  req.ShopId,
	}

	//调用RPC的服务
	adminResp, err := l.svcCtx.AmsRpcClient.PostShopAdmin(l.ctx, &adminReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: adminResp.Code,
		Msg:  adminResp.Msg,
		Data: adminResp.StaffId,
	}, nil
}
