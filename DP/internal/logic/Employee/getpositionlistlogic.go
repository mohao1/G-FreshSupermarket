package Employee

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionListLogic {
	return &GetPositionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetPositionList 获取身份列表
func (l *GetPositionListLogic) GetPositionList(req *types.GetPositionListReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	positionList := bmsclient.GetPositionListReq{
		StaffId: StaffId,
	}

	//调用RPC的服务
	positionListResp, err := l.svcCtx.BmsRpcClient.GetPositionList(l.ctx, &positionList)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "获取信息错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(positionListResp.Code),
		Msg:  positionListResp.Msg,
		Data: positionListResp.Position,
	}, nil
}
