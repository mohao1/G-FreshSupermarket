package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminLogic {
	return &GetAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetAdmin 查看可用管理账号
func (l *GetAdminLogic) GetAdmin(req *types.GetAdminReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	adminReq := amsclient.GetAdminReq{
		AdminId: AdminId,
	}

	//调用RPC的服务
	adminResp, err := l.svcCtx.AmsRpcClient.GetAdmin(l.ctx, &adminReq)
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
		Data: adminResp.ShopAdmin,
	}, nil
}
