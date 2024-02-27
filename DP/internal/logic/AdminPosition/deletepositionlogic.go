package AdminPosition

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePositionLogic {
	return &DeletePositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeletePosition 身份信息删除
func (l *DeletePositionLogic) DeletePosition(req *types.DeletePositionReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	positionReq := amsclient.DeletePositionReq{
		AdminId:    AdminId,
		PositionId: req.PositionId,
	}

	//调用RPC的服务
	positionResp, err := l.svcCtx.AmsRpcClient.DeletePosition(l.ctx, &positionReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: positionResp.Code,
		Msg:  positionResp.Msg,
		Data: positionResp.PositionId,
	}, nil
}
