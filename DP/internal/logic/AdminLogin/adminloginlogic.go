package AdminLogin

import (
	"DP/rpc/Ams/amsclient"
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AdminLogin 系统管理人员登录
func (l *AdminLoginLogic) AdminLogin(req *types.AdminLoginReq) (resp *types.AmsDataResp, err error) {

	//准备数据
	adminLoginReq := amsclient.AdminLoginReq{
		AdminName: req.AdminName,
		PassWord:  req.PassWord,
	}

	//调用RPC的服务
	adminLoginResp, err := l.svcCtx.AmsRpcClient.AdminLogin(l.ctx, &adminLoginReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: adminLoginResp.Code,
		Msg:  adminLoginResp.Msg,
		Data: adminLoginResp.AccessToken,
	}, nil
}
