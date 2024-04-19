package logic

import (
	"context"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSalesRecordsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopSalesRecordsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSalesRecordsListLogic {
	return &GetShopSalesRecordsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 各个店铺商品总销售的数据列表
func (l *GetShopSalesRecordsListLogic) GetShopSalesRecordsList(in *ams.GetShopSalesRecordsListReq) (*ams.GetShopSalesRecordsListResp, error) {
	// todo: add your logic here and delete this line

	return &ams.GetShopSalesRecordsListResp{}, nil
}
