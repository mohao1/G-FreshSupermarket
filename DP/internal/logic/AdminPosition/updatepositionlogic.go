package AdminPosition

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDatePositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDatePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDatePositionLogic {
	return &UpDatePositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpDatePosition 身份信息修改
func (l *UpDatePositionLogic) UpDatePosition(req *types.UpDatePositionReq) (resp *types.AmsDataResp, err error) {
	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	positionReq := amsclient.UpDatePositionReq{
		AdminId:       AdminId,
		PositionId:    req.PositionId,
		PositionName:  req.PositionName,
		PositionGrade: req.PositionGrade,
	}

	//调用RPC的服务
	positionResp, err := l.svcCtx.AmsRpcClient.UpDatePosition(l.ctx, &positionReq)
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
