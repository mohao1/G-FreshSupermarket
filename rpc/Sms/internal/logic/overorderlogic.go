package logic

import (
	"DP/rpc/model"
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type OverOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOverOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OverOrderLogic {
	return &OverOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OverOrder 取消订单
func (l *OverOrderLogic) OverOrder(in *sms.ConfirmOrderReq) (*sms.ConfirmOrderResp, error) {
	orderNumber, err := l.svcCtx.OrderNumberModel.SelectOrderNumberByUserIdAndOrderNumberId(l.ctx, nil, in.UserId, in.OrderNumber)
	if err != nil {
		return nil, err
	}

	//判断是否符合条件
	switch {
	case orderNumber == nil:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "数据错误",
		}, nil
	case orderNumber.OrderOver == 1:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "订单已经取消",
		}, nil
	case orderNumber.OrderOver == 2:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "正在申请推定",
		}, nil
	case orderNumber.ConfirmedDelivery == 1:
		return &sms.ConfirmOrderResp{
			Code:        400,
			OrderNumber: orderNumber.OrderNumber,
			Msg:         "订单已经确认",
		}, nil
	}
	//判断商家是否接单
	if orderNumber.OrderReceive == 0 {
		receive, err2 := l.orderUnReceive(in.UserId, in.OrderNumber)
		if receive {
			return &sms.ConfirmOrderResp{
				Code:        200,
				OrderNumber: orderNumber.OrderNumber,
				Msg:         "订单取消成功",
			}, nil
		} else {
			return nil, err2
		}
	} else {
		receive, err2 := l.orderReceive(in.UserId, in.OrderNumber, orderNumber.ShopId)
		if receive {
			return &sms.ConfirmOrderResp{
				Code:        200,
				OrderNumber: orderNumber.OrderNumber,
				Msg:         "申请退订成功",
			}, nil
		} else {
			return nil, err2
		}
	}

}

// 未确认接单的处理函数
func (l *OverOrderLogic) orderUnReceive(userId string, orderNumber string) (bool, error) {
	err := l.svcCtx.OrderNumberModel.UpDataOrderOver(l.ctx, nil, userId, orderNumber)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 已确认接单的处理函数
func (l *OverOrderLogic) orderReceive(userId string, orderNumber string, shopId string) (bool, error) {
	err := l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err := l.svcCtx.OrderNumberModel.UpDataOrderRefund(ctx, session, userId, orderNumber)
		if err != nil {
			return err
		}

		data := model.RefundRecord{
			RefundId:   uuid.New().String(),
			ShopId:     shopId,
			OrderId:    orderNumber,
			UserId:     userId,
			RefundOver: 0,
			Confirm:    0, //1-同意  2-取消同意
			DeleteKey:  0,
		}
		_, err = l.svcCtx.RefundRecordModel.TransactInsert(ctx, session, &data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
