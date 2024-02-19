package OrderServices

import (
	"DP/rpc/Sms/smsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfo) (resp *types.DataResp, err error) {
	//获取JWT中的UserId
	userId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//调用RPC的服务
	info, err := l.svcCtx.SmsRpcClient.GetUserInfo(l.ctx, &smsclient.UserInfoReq{
		UserId: userId,
	})
	//错误数据处理
	if err != nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "请求错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	//正确数据返回
	return &types.DataResp{
		Code: 200,
		Msg:  "获取成功",
		Data: info,
	}, nil
}
