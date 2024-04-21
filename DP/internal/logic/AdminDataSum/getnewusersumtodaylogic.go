package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNewUserSumToDayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNewUserSumToDayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewUserSumToDayLogic {
	return &GetNewUserSumToDayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetNewUserSumToDay 今日新增用户数量
func (l *GetNewUserSumToDayLogic) GetNewUserSumToDay(req *types.GetNewUserSumToDayReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	newUserSumToDayReq := amsclient.GetNewUserSumToDayReq{
		AdminId: AdminId,
	}

	//调用RPC的服务
	newUserSumToDay, err := l.svcCtx.AmsRpcClient.GetNewUserSumToDay(l.ctx, &newUserSumToDayReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: 200,
		Msg:  "获取成功",
		Data: newUserSumToDay.AddUserSum,
	}, nil
}
