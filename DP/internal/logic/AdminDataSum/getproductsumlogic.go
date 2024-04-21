package AdminDataSum

import (
	"context"

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

func (l *GetProductSumLogic) GetProductSum(req *types.GetProductSumReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
