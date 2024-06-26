// Code generated by goctl. DO NOT EDIT.
// Source: Sms.proto

package server

import (
	"context"

	"DP/rpc/Sms/internal/logic"
	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"
)

type SmsServer struct {
	svcCtx *svc.ServiceContext
	sms.UnimplementedSmsServer
}

func NewSmsServer(svcCtx *svc.ServiceContext) *SmsServer {
	return &SmsServer{
		svcCtx: svcCtx,
	}
}

// 注册
func (s *SmsServer) Register(ctx context.Context, in *sms.RegisterReq) (*sms.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

// 登录
func (s *SmsServer) Login(ctx context.Context, in *sms.LoginReq) (*sms.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// 获取店铺
func (s *SmsServer) GetShop(ctx context.Context, in *sms.GetShopReq) (*sms.GetShopResp, error) {
	l := logic.NewGetShopLogic(ctx, s.svcCtx)
	return l.GetShop(in)
}

// 获取商品列表根据商品类型
func (s *SmsServer) GetProductList(ctx context.Context, in *sms.ProductListReq) (*sms.ProductListResp, error) {
	l := logic.NewGetProductListLogic(ctx, s.svcCtx)
	return l.GetProductList(in)
}

// 获取折扣商品列表根据商品类型
func (s *SmsServer) GetLowProductList(ctx context.Context, in *sms.LowProductListReq) (*sms.LowProductListResp, error) {
	l := logic.NewGetLowProductListLogic(ctx, s.svcCtx)
	return l.GetLowProductList(in)
}

// 获取详细商品信息
func (s *SmsServer) GetDetailedProduct(ctx context.Context, in *sms.DetailedProductReq) (*sms.DetailedProductResp, error) {
	l := logic.NewGetDetailedProductLogic(ctx, s.svcCtx)
	return l.GetDetailedProduct(in)
}

// 计算总价
func (s *SmsServer) GetTotalPrice(ctx context.Context, in *sms.GetTotalPriceReq) (*sms.GetTotalPriceResp, error) {
	l := logic.NewGetTotalPriceLogic(ctx, s.svcCtx)
	return l.GetTotalPrice(in)
}

// 下单
func (s *SmsServer) PostOrder(ctx context.Context, in *sms.PostOrderReq) (*sms.PostOrderResp, error) {
	l := logic.NewPostOrderLogic(ctx, s.svcCtx)
	return l.PostOrder(in)
}

// 获取订单列表
func (s *SmsServer) GetOrderList(ctx context.Context, in *sms.GetOrderListReq) (*sms.GetOrderListResp, error) {
	l := logic.NewGetOrderListLogic(ctx, s.svcCtx)
	return l.GetOrderList(in)
}

// 查看订单详细信息
func (s *SmsServer) GetDetailedOrder(ctx context.Context, in *sms.DetailedOrderReq) (*sms.DetailedOrderResp, error) {
	l := logic.NewGetDetailedOrderLogic(ctx, s.svcCtx)
	return l.GetDetailedOrder(in)
}

// 取消订单
func (s *SmsServer) OverOrder(ctx context.Context, in *sms.ConfirmOrderReq) (*sms.ConfirmOrderResp, error) {
	l := logic.NewOverOrderLogic(ctx, s.svcCtx)
	return l.OverOrder(in)
}

// 取消申请
func (s *SmsServer) CancellationOverOrder(ctx context.Context, in *sms.ConfirmOrderReq) (*sms.ConfirmOrderResp, error) {
	l := logic.NewCancellationOverOrderLogic(ctx, s.svcCtx)
	return l.CancellationOverOrder(in)
}

// 确认订单
func (s *SmsServer) ConfirmOrder(ctx context.Context, in *sms.ConfirmOrderReq) (*sms.ConfirmOrderResp, error) {
	l := logic.NewConfirmOrderLogic(ctx, s.svcCtx)
	return l.ConfirmOrder(in)
}

// 获取对应的店铺的公告列表
func (s *SmsServer) GetAnnouncementList(ctx context.Context, in *sms.AnnouncementListReq) (*sms.AnnouncementListResp, error) {
	l := logic.NewGetAnnouncementListLogic(ctx, s.svcCtx)
	return l.GetAnnouncementList(in)
}

// 获取个人信息
func (s *SmsServer) GetUserInfo(ctx context.Context, in *sms.UserInfoReq) (*sms.UserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}
