package Manage

import (
	"DP/rpc/Bms/bmsclient"
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManageLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManageLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManageLoginLogic {
	return &ManageLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManageLoginLogic) ManageLogin(req *types.ManageLoginReq) (resp *types.BmsDataResp, err error) {

	//构建数据信息
	manageLoginReq := bmsclient.ManageLoginReq{
		StaffId:  req.StaffId,
		Password: req.PassWord,
	}

	//调用RPC的服务
	manageLogin, err := l.svcCtx.BmsRpcClient.ManageLogin(l.ctx, &manageLoginReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "登录错误err:" + err.Error(),
			Data: nil,
		}, err
	}

	return &types.BmsDataResp{
		Code: int(manageLogin.Code),
		Msg:  manageLogin.Msg,
		Data: manageLogin.AccessToken,
	}, nil
}
