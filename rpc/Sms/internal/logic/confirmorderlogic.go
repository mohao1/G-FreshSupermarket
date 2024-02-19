package logic

import (
	"context"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfirmOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfirmOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmOrderLogic {
	return &ConfirmOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ConfirmOrder 确认订单
func (l *ConfirmOrderLogic) ConfirmOrder(in *sms.ConfirmOrderReq) (*sms.ConfirmOrderResp, error) {
	orderNumber, err := l.svcCtx.OrderNumberModel.SelectOrderNumberByUserIdAndOrderNumberId(l.ctx, nil, in.UserId, in.OrderNumber)
	if err != nil {
		return nil, err
	}
	switch {
	case orderNumber.Payment == 0:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "订单还未付款",
		}, nil
	case orderNumber.OrderOver == 1:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "订单已经取消",
		}, nil
	case orderNumber.OrderReceive == 0:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "商家还未接单",
		}, nil
	case orderNumber.ConfirmedDelivery == 1:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "订单已经确认",
		}, nil
	default:
		err2 := l.svcCtx.OrderNumberModel.UpDataOrderConfirm(l.ctx, nil, in.UserId, in.OrderNumber)
		if err2 != nil {
			return nil, err2
		}
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "订单确认完成",
		}, nil
	}
}
