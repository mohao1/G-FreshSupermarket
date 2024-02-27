package AdminUser

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUserList 用户信息列表
func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	userListReq := amsclient.GetUserListReq{
		AdminId: AdminId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	userListResp, err := l.svcCtx.AmsRpcClient.GetUserList(l.ctx, &userListReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: userListResp.Code,
		Msg:  userListResp.Msg,
		Data: userListResp.UserDataList,
	}, nil
}
