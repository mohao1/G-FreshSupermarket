package logic

import (
	"context"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSalesRecordsSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopSalesRecordsSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSalesRecordsSumLogic {
	return &GetShopSalesRecordsSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 销售数据
func (l *GetShopSalesRecordsSumLogic) GetShopSalesRecordsSum(in *ams.GetShopSalesRecordsSumReq) (*ams.GetShopSalesRecordsSumResp, error) {
	// todo: add your logic here and delete this line

	return &ams.GetShopSalesRecordsSumResp{}, nil
}
