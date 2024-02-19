package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type OverOrderHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOverOrderHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OverOrderHandleLogic {
	return &OverOrderHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OverOrderHandle 取消订单申请处理
func (l *OverOrderHandleLogic) OverOrderHandle(in *bms.OverOrderHandleReq) (*bms.OverOrderHandleResp, error) {

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

	//判断是否同意
	if in.Type == 0 {
		err2 := l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			//查询是否存在可以确认取消
			orderNumber, err3 := l.svcCtx.OrderNumberModel.SelectOrderNumberByShopIdAndOrderNumberId(ctx, session, staff.ShopId, in.OrderNumber)
			if err3 != nil {
				return err3
			}
			if orderNumber.OrderOver != 2 {
				return errors.New("申请数据错误")
			}
			//确认取消
			err3 = l.svcCtx.OrderNumberModel.UpDataOrderOver(ctx, session, orderNumber.CustomerId, orderNumber.OrderNumber)
			if err3 != nil {
				return err3
			}

			//取消申请确认
			err3 = l.svcCtx.RefundRecordModel.TransactUpDateOverIsOk(ctx, session, staff.ShopId, in.OrderNumber)
			if err3 != nil {
				return err3
			}

			return nil
		})
		if err2 != nil {
			return nil, err2
		}

		return &bms.OverOrderHandleResp{
			Code:      200,
			Msg:       "同意成功",
			ProductId: in.OrderNumber,
		}, nil
	} else {
		err2 := l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			//查询是否存在可以确认取消
			orderNumber, err3 := l.svcCtx.OrderNumberModel.SelectOrderNumberByShopIdAndOrderNumberId(ctx, session, staff.ShopId, in.OrderNumber)
			if err3 != nil {
				return err3
			}
			if orderNumber.OrderOver != 2 {
				return errors.New("申请数据错误")
			}
			//驳回取消
			err3 = l.svcCtx.OrderNumberModel.UpDataOrderUnRefund(ctx, session, orderNumber.CustomerId, orderNumber.OrderNumber)
			if err3 != nil {
				return err3
			}
			//取消申请驳回
			err3 = l.svcCtx.RefundRecordModel.TransactUpDateOverIsNo(ctx, session, staff.ShopId, in.OrderNumber)
			if err3 != nil {
				return err3
			}

			return nil
		})
		if err2 != nil {
			return nil, err2
		}
		return &bms.OverOrderHandleResp{
			Code:      200,
			Msg:       "驳回成功",
			ProductId: in.OrderNumber,
		}, nil
	}

}
