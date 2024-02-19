package Staff

import (
	"DP/rpc/Bms/bmsclient"
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StaffLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStaffLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StaffLoginLogic {
	return &StaffLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StaffLoginLogic) StaffLogin(req *types.StaffLoginReq) (resp *types.BmsDataResp, err error) {
	//构建数据信息
	staffLoginReq := bmsclient.StaffLoginReq{
		StaffId:  req.StaffId,
		Password: req.PassWord,
	}
	//调用RPC的服务
	loginResp, err := l.svcCtx.BmsRpcClient.StaffLogin(l.ctx, &staffLoginReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "登录错误err:" + err.Error(),
			Data: nil,
		}, nil
	}
	return &types.BmsDataResp{
		Code: int(loginResp.Code),
		Msg:  loginResp.Msg,
		Data: loginResp.AccessToken,
	}, nil
}
