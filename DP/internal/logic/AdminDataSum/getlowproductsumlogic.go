package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLowProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductSumLogic {
	return &GetLowProductSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetLowProductSum 统计折扣商品总量
func (l *GetLowProductSumLogic) GetLowProductSum(req *types.GetLowProductSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	lowProductSumReq := amsclient.GetLowProductSumReq{
		AdminId: AdminId,
	}

	//调用RPC的服务
	lowProductSum, err := l.svcCtx.AmsRpcClient.GetLowProductSum(l.ctx, &lowProductSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: lowProductSum.Code,
		Msg:  lowProductSum.Msg,
		Data: lowProductSum.ProductSum,
	}, nil
}
