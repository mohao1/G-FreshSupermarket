package logic

import (
	"context"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopTimeOrderSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopTimeOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopTimeOrderSumLogic {
	return &GetShopTimeOrderSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 各个店铺根据时间段的订单数量
func (l *GetShopTimeOrderSumLogic) GetShopTimeOrderSum(in *ams.GetShopTimeOrderSumReq) (*ams.GetShopTimeOrderSumResp, error) {
	// todo: add your logic here and delete this line

	return &ams.GetShopTimeOrderSumResp{}, nil
}
