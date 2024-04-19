package logic

import (
	"context"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderSumLogic {
	return &GetOrderSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 今日消费用户数量
func (l *GetOrderSumLogic) GetOrderSum(in *ams.GetOrderSumReq) (*ams.GetOrderSumResp, error) {
	// todo: add your logic here and delete this line

	return &ams.GetOrderSumResp{}, nil
}
