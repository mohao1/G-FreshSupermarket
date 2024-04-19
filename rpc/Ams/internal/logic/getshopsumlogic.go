package logic

import (
	"context"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSumLogic {
	return &GetShopSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 店铺数量
func (l *GetShopSumLogic) GetShopSum(in *ams.GetShopSumReq) (*ams.GetShopSumResp, error) {
	// todo: add your logic here and delete this line

	return &ams.GetShopSumResp{}, nil
}
