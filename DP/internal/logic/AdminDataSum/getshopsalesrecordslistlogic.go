package AdminDataSum

import (
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSalesRecordsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopSalesRecordsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSalesRecordsListLogic {
	return &GetShopSalesRecordsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopSalesRecordsListLogic) GetShopSalesRecordsList(req *types.GetShopSalesRecordsListReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
