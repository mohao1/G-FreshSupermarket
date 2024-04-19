package logic

import (
	"context"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLowProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductSumLogic {
	return &GetLowProductSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 统计折扣商品总量
func (l *GetLowProductSumLogic) GetLowProductSum(in *ams.GetLowProductSumReq) (*ams.GetLowProductSumResp, error) {
	// todo: add your logic here and delete this line

	return &ams.GetLowProductSumResp{}, nil
}
