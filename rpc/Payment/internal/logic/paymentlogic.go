package logic

import (
	"context"

	"DP/rpc/Payment/internal/svc"
	"DP/rpc/Payment/pb/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentLogic {
	return &PaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Payment 付款
func (l *PaymentLogic) Payment(in *payment.PaymentReq) (*payment.PaymentResp, error) {
	// todo: add your logic here and delete this line

	return &payment.PaymentResp{}, nil
}
