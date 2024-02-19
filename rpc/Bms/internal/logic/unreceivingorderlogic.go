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

type UnReceivingOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewUnReceivingOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnReceivingOrderLogic {
	return &UnReceivingOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UnReceivingOrder 取消接单
func (l *UnReceivingOrderLogic) UnReceivingOrder(in *bms.UnReceivingOrderReq) (*bms.UnReceivingOrderResp, error) {
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

		//判断订单是否未被接收
		if orderNumber.OrderReceive != 1 {
			return errors.New("订单未被接收")
		}

		//进行取消接单操作
		err2 = l.svcCtx.OrderNumberModel.UpDateOrderReceive(ctx, session, 0, in.OrderNumber)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &bms.UnReceivingOrderResp{
		Code:        200,
		Msg:         "取消接单成功",
		OrderNumber: in.OrderNumber,
	}, nil
}
