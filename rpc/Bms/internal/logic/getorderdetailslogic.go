package logic

import (
	"context"
	"errors"
	"strconv"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailsLogic {
	return &GetOrderDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderDetails 查看订单详情
func (l *GetOrderDetailsLogic) GetOrderDetails(in *bms.OrderDetailsReq) (*bms.OrderDetailsResp, error) {
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//查询对应订单
	orderNumber, err := l.svcCtx.OrderNumberModel.FindOne(l.ctx, in.OrderNumber)
	if err != nil {
		return nil, err
	}

	//查询Shop店铺
	shop, err := l.svcCtx.ShopModel.FindOne(l.ctx, staff.ShopId)
	if err != nil {
		return nil, err
	}

	//进行店铺和订单的匹配判断
	if orderNumber.ShopId != staff.ShopId {
		return nil, errors.New("店铺和订单不匹配")
	}

	orderList, err := l.svcCtx.OrderModel.SelectOrderByOrderNumberId(l.ctx, orderNumber.OrderNumber)
	if err != nil {
		return nil, err
	}

	orderProductList := make([]*bms.OrderProduct, len(*orderList))

	for i, order := range *orderList {

		//根据类型Id查询货物类型信息
		productType, err2 := l.svcCtx.ProductTypeModel.FindOne(l.ctx, order.ProductTypeId)
		if err2 != nil {
			return nil, err2
		}

		orderProductList[i] = &bms.OrderProduct{
			OrderName:       order.OrderName,
			OrderTitle:      order.OrderTitle,
			Price:           order.Price,
			ProductQuantity: order.OrderQuantity,
			ProductSize:     order.ProductSize,
			ProductType:     productType.ProductTypeName,
			ProductUnit:     productType.ProductUnit,
			ProductPicture:  order.ProductPicture,
		}
	}

	OrderDetailsResp := bms.OrderDetailsResp{
		OrderNumber:       orderNumber.OrderNumber,
		ShopName:          shop.ShopName,
		Total:             strconv.Itoa(int(orderNumber.Total)),
		TotalPrice:        orderNumber.TotalPrice,
		OrderProduct:      orderProductList,
		Payment:           strconv.Itoa(int(orderNumber.Payment)),
		OrderOver:         strconv.Itoa(int(orderNumber.OrderOver)),
		ConfirmedDelivery: strconv.Itoa(int(orderNumber.ConfirmedDelivery)),
		OrderReceive:      strconv.Itoa(int(orderNumber.OrderReceive)),
		Notes:             orderNumber.Notes,
		CreationTime:      orderNumber.CreationTime.String(),
		ConfirmTime:       "",
		DeliveryTime:      orderNumber.DeliveryTime.String(),
	}

	if orderNumber.ConfirmTime.Valid {
		OrderDetailsResp.ConfirmTime = orderNumber.ConfirmTime.Time.String()
	}

	return &OrderDetailsResp, nil
}
