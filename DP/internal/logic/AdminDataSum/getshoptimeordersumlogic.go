package AdminDataSum

import (
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopTimeOrderSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopTimeOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopTimeOrderSumLogic {
	return &GetShopTimeOrderSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopTimeOrderSumLogic) GetShopTimeOrderSum(req *types.GetShopTimeOrderSumReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
