package Employee

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetEmployeeLogic {
	return &SetEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SetEmployee 设置员工
func (l *SetEmployeeLogic) SetEmployee(req *types.SetEmployeeReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	employee := bmsclient.SetEmployeeReq{
		StaffId:    StaffId,
		NewStaffId: req.NewStaffId,
		StaffName:  req.StaffName,
		PositionId: req.PositionId,
		PassWord:   req.PassWord,
	}

	//调用RPC的服务
	employeeResp, err := l.svcCtx.BmsRpcClient.SetEmployee(l.ctx, &employee)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "获取信息错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(employeeResp.Code),
		Msg:  employeeResp.Msg,
		Data: employeeResp.StaffId,
	}, nil
}
