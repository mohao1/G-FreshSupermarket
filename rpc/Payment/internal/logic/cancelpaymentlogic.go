package logic

import (
	"context"

	"DP/rpc/Payment/internal/svc"
	"DP/rpc/Payment/pb/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelPaymentLogic {
	return &CancelPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CancelPayment 取消付款
func (l *CancelPaymentLogic) CancelPayment(in *payment.CancelPaymentReq) (*payment.CancelPaymentResp, error) {
	// todo: add your logic here and delete this line

	return &payment.CancelPaymentResp{}, nil
}
