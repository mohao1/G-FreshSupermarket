package logic

import (
	"context"
	"fmt"
	"strconv"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnreceivedOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUnreceivedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnreceivedOrderLogic {
	return &GetUnreceivedOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUnreceivedOrder 查看未接订单
func (l *GetUnreceivedOrderLogic) GetUnreceivedOrder(in *bms.UnreceivedOrderReq) (*bms.UnreceivedOrderResp, error) {

	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//查询Shop店铺
	shop, err := l.svcCtx.ShopModel.FindOne(l.ctx, staff.ShopId)
	if err != nil {
		return nil, err
	}

	//查询没有确认了的订单
	orders, err := l.svcCtx.OrderNumberModel.SelectReceivedOrderListByShop(l.ctx, nil, staff.ShopId, 0, 0, int(in.Limit))
	if err != nil {
		fmt.Println("--------")
		return nil, err
	}

	orderList := make([]*bms.Order, len(*orders))

	for i, order := range *orders {

		//查询用户名称
		user, err2 := l.svcCtx.UserModel.FindOne(l.ctx, order.CustomerId)
		if err2 != nil {
			return nil, err2
		}

		orderList[i] = &bms.Order{
			OrderNumber:       order.OrderNumber,
			CustomerName:      user.Name,
			ShopName:          shop.ShopName,
			Total:             strconv.Itoa(int(order.Total)),
			TotalPrice:        order.TotalPrice,
			Payment:           strconv.Itoa(int(order.Payment)),
			OrderOver:         order.OrderOver,
			OrderReceive:      order.OrderReceive,
			ConfirmedDelivery: order.ConfirmedDelivery,
			DeliveryTime:      order.DeliveryTime.String(),
		}
	}

	return &bms.UnreceivedOrderResp{
		OrderList: orderList,
	}, nil
}
