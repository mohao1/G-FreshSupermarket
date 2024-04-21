package AdminDataSum

import (
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopOrderSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopOrderSumLogic {
	return &GetShopOrderSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopOrderSumLogic) GetShopOrderSum(req *types.GetShopOrderSumReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
