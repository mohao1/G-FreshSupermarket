package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"sync"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceivingOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewReceivingOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceivingOrderLogic {
	return &ReceivingOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReceivingOrder 接单
func (l *ReceivingOrderLogic) ReceivingOrder(in *bms.ReceivingOrderReq) (*bms.ReceivingOrderResp, error) {

	//查询店铺信息和个人的信息
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//上锁
	l.rWMutex.Lock()
	defer l.rWMutex.Unlock()

	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//查询订单是否存在
		orderNumber, err2 := l.svcCtx.OrderNumberModel.SelectOrderNumberByOrderId(ctx, session, in.OrderNumber)
		if err2 != nil {
			return err2
		}

		//匹配店铺信息
		if staff.ShopId != orderNumber.ShopId {
			return errors.New("店铺信息错误")
		}

		//判断订单是否已被接收
		if orderNumber.OrderReceive != 0 {
			return errors.New("订单已被接收")
		}

		//进行接单操作
		err2 = l.svcCtx.OrderNumberModel.UpDateOrderReceive(ctx, session, 1, in.OrderNumber)
		if err2 != nil {
			return err2
		}

		return nil

	})
	if err != nil {
		return nil, err
	}

	return &bms.ReceivingOrderResp{
		Code:        200,
		Msg:         "接单完成",
		OrderNumber: in.OrderNumber,
	}, nil
}
