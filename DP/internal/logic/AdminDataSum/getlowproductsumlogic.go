package AdminDataSum

import (
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLowProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductSumLogic {
	return &GetLowProductSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLowProductSumLogic) GetLowProductSum(req *types.GetLowProductSumReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
