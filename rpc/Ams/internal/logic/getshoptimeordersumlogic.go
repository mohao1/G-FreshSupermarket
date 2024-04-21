package logic

import (
	"context"
	"errors"
	"time"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopTimeOrderSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopTimeOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopTimeOrderSumLogic {
	return &GetShopTimeOrderSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopTimeOrderSum 各个店铺根据时间段的订单数量
func (l *GetShopTimeOrderSumLogic) GetShopTimeOrderSum(in *ams.GetShopTimeOrderSumReq) (*ams.GetShopTimeOrderSumResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	orderSum, err := l.svcCtx.OrderNumberModel.SelectOrderTheStartTimeAndEndTimeByShopIdSum(l.ctx, in.ShopId, time.Unix(in.TopTime, 0), time.Unix(in.GetEndTime(), 0))
	if err != nil {
		return nil, err
	}

	shop, err := l.svcCtx.ShopModel.FindOne(l.ctx, in.ShopId)
	if err != nil {
		return nil, err
	}

	list := make([]*ams.ShopTimeOrderData, len(*orderSum))

	for i, order := range *orderSum {
		list[i] = &ams.ShopTimeOrderData{
			Time:     order.Time.String(),
			ShopId:   in.ShopId,
			ShopName: shop.ShopName,
			OrderSum: order.OrderCount,
		}
	}

	return &ams.GetShopTimeOrderSumResp{
		Code:              200,
		Msg:               "获取成功",
		ShopTimeOrderList: list,
	}, nil
}
