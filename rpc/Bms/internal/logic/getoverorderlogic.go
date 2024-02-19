package logic

import (
	"context"
	"errors"
	"strconv"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOverOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOverOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOverOrderLogic {
	return &GetOverOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOverOrder 获取申请取消的订单列表
func (l *GetOverOrderLogic) GetOverOrder(in *bms.OverOrderListReq) (*bms.OverOrderListResp, error) {

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

	shop, err := l.svcCtx.ShopModel.FindOne(l.ctx, staff.ShopId)
	if err != nil {
		return nil, err
	}

	refundRecords, err := l.svcCtx.RefundRecordModel.SelectRefundRecordListByShop(l.ctx, staff.ShopId, 0, 0, int(in.Limit))
	if err != nil {
		return nil, err
	}

	//创建列表
	refundRecordList := make([]*bms.Order, len(*refundRecords))

	//填充列表
	for i, refundRecord := range *refundRecords {

		//查询订单
		orderNumber, err2 := l.svcCtx.OrderNumberModel.FindOne(l.ctx, refundRecord.OrderId)
		if err2 != nil {
			return nil, err2
		}

		//查询用户信息
		user, err2 := l.svcCtx.UserModel.FindOne(l.ctx, orderNumber.CustomerId)
		if err2 != nil {
			return nil, err2
		}

		refundRecordList[i] = &bms.Order{
			OrderNumber:       orderNumber.OrderNumber,
			CustomerName:      user.Name,
			ShopName:          shop.ShopName,
			Total:             strconv.Itoa(int(orderNumber.Total)),
			TotalPrice:        orderNumber.TotalPrice,
			Payment:           strconv.Itoa(int(orderNumber.Payment)),
			OrderOver:         orderNumber.OrderOver,
			OrderReceive:      orderNumber.OrderReceive,
			ConfirmedDelivery: orderNumber.ConfirmedDelivery,
			DeliveryTime:      orderNumber.DeliveryTime.String(),
		}
	}

	return &bms.OverOrderListResp{
		Code:      200,
		Msg:       "获取成功",
		OrderList: refundRecordList,
	}, nil
}
