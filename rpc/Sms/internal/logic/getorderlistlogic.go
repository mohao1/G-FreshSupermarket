package logic

import (
	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderList 获取订单列表
func (l *GetOrderListLogic) GetOrderList(in *sms.GetOrderListReq) (*sms.GetOrderListResp, error) {
	orderNumber, err := l.svcCtx.OrderNumberModel.SelectOrderNumber(l.ctx, in.UserId, in.Limit)
	if err != nil {
		return nil, err
	}

	orderList := make([]*sms.GetOrder, len(*orderNumber))
	for i, order := range *orderNumber {
		findOne, _ := l.svcCtx.ShopModel.FindOne(l.ctx, order.ShopId)
		orderList[i] = &sms.GetOrder{
			OrderNumber:       order.OrderNumber,
			ShopName:          findOne.ShopId,
			Total:             strconv.Itoa(int(order.Total)),
			TotalPrice:        order.TotalPrice,
			Payment:           strconv.Itoa(int(order.Payment)),
			OrderOver:         strconv.Itoa(int(order.OrderOver)),
			OrderReceive:      strconv.Itoa(int(order.OrderReceive)),
			ConfirmedDelivery: strconv.Itoa(int(order.ConfirmedDelivery)),
			CreationTime:      order.CreationTime.String(),
		}
	}
	return &sms.GetOrderListResp{Order: orderList}, nil
}
