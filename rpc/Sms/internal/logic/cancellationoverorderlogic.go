package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancellationOverOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancellationOverOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancellationOverOrderLogic {
	return &CancellationOverOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CancellationOverOrder 取消申请
func (l *CancellationOverOrderLogic) CancellationOverOrder(in *sms.ConfirmOrderReq) (*sms.ConfirmOrderResp, error) {
	orderNumber, err := l.svcCtx.OrderNumberModel.SelectOrderNumberByUserIdAndOrderNumberId(l.ctx, nil, in.UserId, in.OrderNumber)
	if err != nil {
		return nil, err
	}
	switch {
	case orderNumber == nil:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "信息错误",
		}, nil
	case orderNumber.OrderReceive == 0:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "订单没有确认接单",
		}, nil
	case orderNumber.OrderOver == 0:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "订单没有申请",
		}, nil
	case orderNumber.OrderOver == 1:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "订单已经取消",
		}, nil
	}

	if orderNumber.OrderOver == 2 {
		err2 := l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			res, err3 := l.svcCtx.OrderNumberModel.SelectOrderNumberByUserIdAndOrderNumberId(l.ctx, nil, in.UserId, in.OrderNumber)
			if err3 != nil {
				return err3
			}
			if res.OrderOver != 2 {
				return errors.New("数据已经变化")
			}
			err3 = l.svcCtx.OrderNumberModel.UpDataOrderUnRefund(ctx, session, in.UserId, in.OrderNumber)
			if err3 != nil {
				return err3
			}
			err4 := l.svcCtx.RefundRecordModel.TransactUpDataRefundOver(ctx, session, in.OrderNumber, in.UserId)
			if err4 != nil {
				return err4
			}
			return nil
		})
		if err2 != nil {
			return nil, err2
		}
	} else {
		return &sms.ConfirmOrderResp{
			Code:        200,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "数据错误",
		}, nil
	}

	return &sms.ConfirmOrderResp{
		Code:        200,
		OrderNumber: orderNumber.OrderNumber,
		Msg:         "取消成功",
	}, nil
}
