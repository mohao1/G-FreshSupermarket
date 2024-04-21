package AdminDataSum

import (
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSalesRecordsSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopSalesRecordsSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSalesRecordsSumLogic {
	return &GetShopSalesRecordsSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopSalesRecordsSumLogic) GetShopSalesRecordsSum(req *types.GetShopSalesRecordsSumReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
