package logic

import (
	"context"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSumLogic {
	return &GetProductSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 统计普通商品总量
func (l *GetProductSumLogic) GetProductSum(in *ams.GetProductSumReq) (*ams.GetProductSumResp, error) {
	// todo: add your logic here and delete this line

	return &ams.GetProductSumResp{}, nil
}
