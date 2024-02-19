package logic

import (
	"context"
	"errors"
	"strconv"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailedOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDetailedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailedOrderLogic {
	return &GetDetailedOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetDetailedOrder 查看订单详细信息
func (l *GetDetailedOrderLogic) GetDetailedOrder(in *sms.DetailedOrderReq) (*sms.DetailedOrderResp, error) {

	//查询对应的订单信息
	orderNumber, err := l.svcCtx.OrderNumberModel.SelectOrderNumberByUserIdAndOrderNumberId(l.ctx, nil, in.UserId, in.OrderNumber)
	if err != nil {
		return nil, err
	}
	if orderNumber == nil {
		return nil, errors.New("数据错误")
	}

	//查询对应订单号的订单信息列表
	order, err := l.svcCtx.OrderModel.SelectOrderByOrderNumberId(l.ctx, in.OrderNumber)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("数据错误")
	}

	one, err := l.svcCtx.ShopModel.FindOne(l.ctx, orderNumber.ShopId)
	if err != nil {
		return nil, err
	}
	if one == nil {
		return nil, errors.New("数据错误")
	}

	data := sms.DetailedOrderResp{
		OrderNumber:       orderNumber.OrderNumber,
		ShopName:          one.ShopName,
		Total:             strconv.Itoa(int(orderNumber.Total)),
		TotalPrice:        orderNumber.TotalPrice,
		OrderProduct:      nil,
		Payment:           strconv.Itoa(int(orderNumber.Payment)),
		OrderOver:         strconv.Itoa(int(orderNumber.OrderOver)),
		ConfirmedDelivery: strconv.Itoa(int(orderNumber.ConfirmedDelivery)),
		OrderReceive:      strconv.Itoa(int(orderNumber.OrderReceive)),
		Notes:             orderNumber.Notes,
		CreationTime:      orderNumber.CreationTime.String(),
		ConfirmTime:       "",
		DeliveryTime:      orderNumber.DeliveryTime.String(),
	}
	OrderProduct := make([]*sms.OrderProduct, len(*order))

	if orderNumber.ConfirmTime.Valid {
		data.ConfirmTime = orderNumber.ConfirmTime.Time.String()
	}

	for i, orderItem := range *order {
		findOne, err2 := l.svcCtx.ProductType.FindOne(l.ctx, orderItem.ProductTypeId)
		if err2 != nil {
			return nil, err2
		}
		if findOne == nil {
			return nil, errors.New("数据错误")
		}

		OrderProduct[i] = &sms.OrderProduct{
			OrderName:       orderItem.OrderName,
			OrderTitle:      orderItem.OrderTitle,
			Price:           orderItem.Price,
			ProductQuantity: orderItem.OrderQuantity,
			ProductSize:     orderItem.ProductSize,
			ProductType:     findOne.ProductTypeName,
			ProductUnit:     findOne.ProductUnit,
			ProductPicture:  orderItem.ProductPicture,
		}
	}
	data.OrderProduct = OrderProduct
	return &data, nil
}
