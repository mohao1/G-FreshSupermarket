package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSumLogic {
	return &GetUserSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUserSum 用户人数
func (l *GetUserSumLogic) GetUserSum(req *types.GetUserSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	userSumReq := amsclient.GetUserSumReq{
		AdminId: AdminId,
	}

	//调用RPC的服务
	userSum, err := l.svcCtx.AmsRpcClient.GetUserSum(l.ctx, &userSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: userSum.Code,
		Msg:  userSum.Msg,
		Data: userSum.UserSum,
	}, nil
}
