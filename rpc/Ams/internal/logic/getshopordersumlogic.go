package logic

import (
	"context"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopOrderSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopOrderSumLogic {
	return &GetShopOrderSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopOrderSum 各个店铺总的订单数量
func (l *GetShopOrderSumLogic) GetShopOrderSum(in *ams.GetShopOrderSumReq) (*ams.GetShopOrderSumResp, error) {
	// todo: add your logic here and delete this line

	return &ams.GetShopOrderSumResp{}, nil
}
