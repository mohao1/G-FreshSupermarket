package Employee

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmployeeLogic {
	return &DeleteEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEmployeeLogic) DeleteEmployee(req *types.DeleteEmployeeReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	deleteEmployeeReq := bmsclient.DeleteEmployeeReq{
		StaffId:    StaffId,
		SetStaffId: req.DeleteStaffId,
	}

	//调用RPC的服务
	deleteEmployeeResp, err := l.svcCtx.BmsRpcClient.DeleteEmployee(l.ctx, &deleteEmployeeReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "获取信息错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(deleteEmployeeResp.Code),
		Msg:  deleteEmployeeResp.Msg,
		Data: deleteEmployeeResp.StaffId,
	}, nil
}
