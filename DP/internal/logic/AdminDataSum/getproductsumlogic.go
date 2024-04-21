package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSumLogic {
	return &GetProductSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetProductSum 统计普通商品总量
func (l *GetProductSumLogic) GetProductSum(req *types.GetProductSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	productSumReq := amsclient.GetProductSumReq{
		AdminId: AdminId,
	}

	//调用RPC的服务
	productSum, err := l.svcCtx.AmsRpcClient.GetProductSum(l.ctx, &productSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: productSum.Code,
		Msg:  productSum.Msg,
		Data: productSum.ProductSum,
	}, nil
}
