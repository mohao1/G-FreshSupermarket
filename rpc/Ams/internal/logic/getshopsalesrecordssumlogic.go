package logic

import (
	"context"
	"errors"

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

// GetShopSalesRecordsSum 销售数据
func (l *GetShopSalesRecordsSumLogic) GetShopSalesRecordsSum(in *ams.GetShopSalesRecordsSumReq) (*ams.GetShopSalesRecordsSumResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	shopSalesRecordsSum, err := l.svcCtx.OrderModel.SelectShopSalesRecordsSum(l.ctx, in.ShopId)
	if err != nil {
		return nil, err
	}

	return &ams.GetShopSalesRecordsSumResp{
		Code:                200,
		Msg:                 "获取成功",
		ShopSalesRecordsSum: shopSalesRecordsSum.ProductSalesRecordsSum,
	}, nil
}
