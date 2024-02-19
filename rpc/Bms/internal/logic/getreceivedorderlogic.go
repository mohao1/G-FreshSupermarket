package logic

import (
	"context"
	"strconv"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReceivedOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReceivedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReceivedOrderLogic {
	return &GetReceivedOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetReceivedOrder 查看已接订单
func (l *GetReceivedOrderLogic) GetReceivedOrder(in *bms.ReceivedOrderReq) (*bms.ReceivedOrderResp, error) {

	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//查询Shop店铺
	shop, err := l.svcCtx.ShopModel.FindOne(l.ctx, staff.ShopId)
	if err != nil {
		return nil, err
	}

	//查询已经确认了的订单
	orders, err := l.svcCtx.OrderNumberModel.SelectReceivedOrderListByShop(l.ctx, nil, staff.ShopId, 1, 0, int(in.Limit))
	if err != nil {
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

	return &bms.ReceivedOrderResp{
		OrderList: orderList,
	}, nil
}
