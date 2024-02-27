package AdminPosition

import (
	"DP/rpc/Ams/amsclient"
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

// GetPositionList 身份信息列表查看
func (l *GetPositionListLogic) GetPositionList(req *types.GetPositionListAdminReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	positionListReq := amsclient.GetPositionListReq{
		AdminId: AdminId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	positionListResp, err := l.svcCtx.AmsRpcClient.GetPositionList(l.ctx, &positionListReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: positionListResp.Code,
		Msg:  positionListResp.Msg,
		Data: positionListResp.PositionDataList,
	}, nil
}
