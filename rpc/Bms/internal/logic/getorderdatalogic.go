package logic

import (
	"context"
	"errors"
	"strconv"
	"time"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDataLogic {
	return &GetOrderDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderData 查看订单数据
func (l *GetOrderDataLogic) GetOrderData(in *bms.OrderDataReq) (*bms.OrderDataResp, error) {

	//查询店铺信息和个人的信息
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//进行权限判断
	position, err := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
	if err != nil {
		return nil, err
	}
	if position.PositionName != "经理" {
		return nil, errors.New("权限不足")
	}

	OrderDataList, err := l.svcCtx.OrderNumberModel.SelectOrderDataListByShopAndTime(l.ctx, staff.ShopId, time.Unix(in.GetTime, 0))
	if err != nil {
		return nil, err
	}

	var totalPrice int64 = 0
	var total int64 = 0

	//进行数据计算整合
	for _, orderNumber := range *OrderDataList {
		total = total + orderNumber.Total
		atoi, err2 := strconv.Atoi(orderNumber.TotalPrice)
		if err2 != nil {
			return nil, err2
		}
		totalPrice = totalPrice + int64(atoi)
	}

	OrderData := &bms.OrderData{
		OrderQuantity: int64(len(*OrderDataList)),
		TotalPrice:    totalPrice,
		Total:         total,
	}

	return &bms.OrderDataResp{
		Code:      200,
		Msg:       "获取成功",
		OrderData: OrderData,
	}, nil
}
