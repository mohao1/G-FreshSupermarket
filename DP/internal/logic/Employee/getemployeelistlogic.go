package Employee

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmployeeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmployeeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmployeeListLogic {
	return &GetEmployeeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetEmployeeList 获取员工列表
func (l *GetEmployeeListLogic) GetEmployeeList(req *types.GetEmployeeListReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	getEmployeeListReq := bmsclient.GetEmployeeListReq{
		StaffId: StaffId,
	}

	//调用RPC的服务
	getEmployeeListResp, err := l.svcCtx.BmsRpcClient.GetEmployeeList(l.ctx, &getEmployeeListReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "获取信息错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(getEmployeeListResp.Code),
		Msg:  getEmployeeListResp.Msg,
		Data: getEmployeeListResp.StaffList,
	}, nil
}
