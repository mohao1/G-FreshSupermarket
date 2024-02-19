package Employee

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateEmployeeLogic {
	return &UpDateEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpDateEmployeeLogic) UpDateEmployee(req *types.UpDateEmployeeReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	upDateEmployeeReq := bmsclient.UpDateEmployeeReq{
		StaffId:    StaffId,
		SetStaffId: req.UpdateStaffId,
		StaffName:  req.StaffName,
		PositionId: req.PositionId,
		PassWord:   req.PassWord,
	}

	//调用RPC的服务

	upDateEmployeeResp, err := l.svcCtx.BmsRpcClient.UpDateEmployee(l.ctx, &upDateEmployeeReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "获取信息错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(upDateEmployeeResp.Code),
		Msg:  upDateEmployeeResp.Msg,
		Data: upDateEmployeeResp.StaffId,
	}, nil
}
