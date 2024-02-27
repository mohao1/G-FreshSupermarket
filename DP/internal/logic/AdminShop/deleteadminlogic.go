package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAdminLogic {
	return &DeleteAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteAdmin 删除可用管理账号
func (l *DeleteAdminLogic) DeleteAdmin(req *types.DeleteAdminReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	adminReq := amsclient.DeleteAdminReq{
		AdminId: AdminId,
		StaffId: req.StaffId,
	}

	//调用RPC的服务
	adminResp, err := l.svcCtx.AmsRpcClient.DeleteAdmin(l.ctx, &adminReq)
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
