package AdminPosition

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostPositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostPositionLogic {
	return &PostPositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PostPosition 身份信息设置
func (l *PostPositionLogic) PostPosition(req *types.PostPositionReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	positionReq := amsclient.PostPositionReq{
		AdminId:       AdminId,
		PositionName:  req.PositionName,
		PositionGrade: req.PositionGrade,
	}

	//调用RPC的服务
	positionResp, err := l.svcCtx.AmsRpcClient.PostPosition(l.ctx, &positionReq)
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
