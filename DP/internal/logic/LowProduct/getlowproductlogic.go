package LowProduct

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLowProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductLogic {
	return &GetLowProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetLowProduct 查看普通商品列表
func (l *GetLowProductLogic) GetLowProduct(req *types.GetLowProductReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	lowProductReq := bmsclient.GetLowProductReq{
		StaffId: StaffId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	lowProductResp, err := l.svcCtx.BmsRpcClient.GetLowProduct(l.ctx, &lowProductReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(lowProductResp.Code),
		Msg:  lowProductResp.Msg,
		Data: lowProductResp.LowProduct,
	}, nil
}
