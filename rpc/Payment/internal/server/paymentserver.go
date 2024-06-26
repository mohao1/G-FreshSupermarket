// Code generated by goctl. DO NOT EDIT.
// Source: Payment.proto

package server

import (
	"context"

	"DP/rpc/Payment/internal/logic"
	"DP/rpc/Payment/internal/svc"
	"DP/rpc/Payment/pb/payment"
)

type PaymentServer struct {
	svcCtx *svc.ServiceContext
	payment.UnimplementedPaymentServer
}

func NewPaymentServer(svcCtx *svc.ServiceContext) *PaymentServer {
	return &PaymentServer{
		svcCtx: svcCtx,
	}
}

// 付款
func (s *PaymentServer) Payment(ctx context.Context, in *payment.PaymentReq) (*payment.PaymentResp, error) {
	l := logic.NewPaymentLogic(ctx, s.svcCtx)
	return l.Payment(in)
}

// 取消付款
func (s *PaymentServer) CancelPayment(ctx context.Context, in *payment.CancelPaymentReq) (*payment.CancelPaymentResp, error) {
	l := logic.NewCancelPaymentLogic(ctx, s.svcCtx)
	return l.CancelPayment(in)
}
