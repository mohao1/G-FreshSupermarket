// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: Sms.proto

package sms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Sms_Register_FullMethodName              = "/sms.Sms/Register"
	Sms_Login_FullMethodName                 = "/sms.Sms/Login"
	Sms_GetShop_FullMethodName               = "/sms.Sms/GetShop"
	Sms_GetProductList_FullMethodName        = "/sms.Sms/GetProductList"
	Sms_GetLowProductList_FullMethodName     = "/sms.Sms/GetLowProductList"
	Sms_GetDetailedProduct_FullMethodName    = "/sms.Sms/GetDetailedProduct"
	Sms_GetTotalPrice_FullMethodName         = "/sms.Sms/GetTotalPrice"
	Sms_PostOrder_FullMethodName             = "/sms.Sms/PostOrder"
	Sms_GetOrderList_FullMethodName          = "/sms.Sms/GetOrderList"
	Sms_GetDetailedOrder_FullMethodName      = "/sms.Sms/GetDetailedOrder"
	Sms_OverOrder_FullMethodName             = "/sms.Sms/OverOrder"
	Sms_CancellationOverOrder_FullMethodName = "/sms.Sms/CancellationOverOrder"
	Sms_ConfirmOrder_FullMethodName          = "/sms.Sms/ConfirmOrder"
	Sms_GetAnnouncementList_FullMethodName   = "/sms.Sms/GetAnnouncementList"
	Sms_GetUserInfo_FullMethodName           = "/sms.Sms/GetUserInfo"
)

// SmsClient is the client API for Sms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsClient interface {
	// 注册
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	// 登录
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	// 获取店铺
	GetShop(ctx context.Context, in *GetShopReq, opts ...grpc.CallOption) (*GetShopResp, error)
	// 获取商品列表根据商品类型
	GetProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListResp, error)
	// 获取折扣商品列表根据商品类型
	GetLowProductList(ctx context.Context, in *LowProductListReq, opts ...grpc.CallOption) (*LowProductListResp, error)
	// 获取详细商品信息
	GetDetailedProduct(ctx context.Context, in *DetailedProductReq, opts ...grpc.CallOption) (*DetailedProductResp, error)
	// 计算总价
	GetTotalPrice(ctx context.Context, in *GetTotalPriceReq, opts ...grpc.CallOption) (*GetTotalPriceResp, error)
	// 下单
	PostOrder(ctx context.Context, in *PostOrderReq, opts ...grpc.CallOption) (*PostOrderResp, error)
	// 获取订单列表
	GetOrderList(ctx context.Context, in *GetOrderListReq, opts ...grpc.CallOption) (*GetOrderListResp, error)
	// 查看订单详细信息
	GetDetailedOrder(ctx context.Context, in *DetailedOrderReq, opts ...grpc.CallOption) (*DetailedOrderResp, error)
	// 取消订单
	OverOrder(ctx context.Context, in *ConfirmOrderReq, opts ...grpc.CallOption) (*ConfirmOrderResp, error)
	// 取消申请
	CancellationOverOrder(ctx context.Context, in *ConfirmOrderReq, opts ...grpc.CallOption) (*ConfirmOrderResp, error)
	// 确认订单
	ConfirmOrder(ctx context.Context, in *ConfirmOrderReq, opts ...grpc.CallOption) (*ConfirmOrderResp, error)
	// 获取对应的店铺的公告列表
	GetAnnouncementList(ctx context.Context, in *AnnouncementListReq, opts ...grpc.CallOption) (*AnnouncementListResp, error)
	// 获取个人信息
	GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
}

type smsClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsClient(cc grpc.ClientConnInterface) SmsClient {
	return &smsClient{cc}
}

func (c *smsClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, Sms_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Sms_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetShop(ctx context.Context, in *GetShopReq, opts ...grpc.CallOption) (*GetShopResp, error) {
	out := new(GetShopResp)
	err := c.cc.Invoke(ctx, Sms_GetShop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListResp, error) {
	out := new(ProductListResp)
	err := c.cc.Invoke(ctx, Sms_GetProductList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetLowProductList(ctx context.Context, in *LowProductListReq, opts ...grpc.CallOption) (*LowProductListResp, error) {
	out := new(LowProductListResp)
	err := c.cc.Invoke(ctx, Sms_GetLowProductList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetDetailedProduct(ctx context.Context, in *DetailedProductReq, opts ...grpc.CallOption) (*DetailedProductResp, error) {
	out := new(DetailedProductResp)
	err := c.cc.Invoke(ctx, Sms_GetDetailedProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetTotalPrice(ctx context.Context, in *GetTotalPriceReq, opts ...grpc.CallOption) (*GetTotalPriceResp, error) {
	out := new(GetTotalPriceResp)
	err := c.cc.Invoke(ctx, Sms_GetTotalPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) PostOrder(ctx context.Context, in *PostOrderReq, opts ...grpc.CallOption) (*PostOrderResp, error) {
	out := new(PostOrderResp)
	err := c.cc.Invoke(ctx, Sms_PostOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetOrderList(ctx context.Context, in *GetOrderListReq, opts ...grpc.CallOption) (*GetOrderListResp, error) {
	out := new(GetOrderListResp)
	err := c.cc.Invoke(ctx, Sms_GetOrderList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetDetailedOrder(ctx context.Context, in *DetailedOrderReq, opts ...grpc.CallOption) (*DetailedOrderResp, error) {
	out := new(DetailedOrderResp)
	err := c.cc.Invoke(ctx, Sms_GetDetailedOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) OverOrder(ctx context.Context, in *ConfirmOrderReq, opts ...grpc.CallOption) (*ConfirmOrderResp, error) {
	out := new(ConfirmOrderResp)
	err := c.cc.Invoke(ctx, Sms_OverOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) CancellationOverOrder(ctx context.Context, in *ConfirmOrderReq, opts ...grpc.CallOption) (*ConfirmOrderResp, error) {
	out := new(ConfirmOrderResp)
	err := c.cc.Invoke(ctx, Sms_CancellationOverOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) ConfirmOrder(ctx context.Context, in *ConfirmOrderReq, opts ...grpc.CallOption) (*ConfirmOrderResp, error) {
	out := new(ConfirmOrderResp)
	err := c.cc.Invoke(ctx, Sms_ConfirmOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetAnnouncementList(ctx context.Context, in *AnnouncementListReq, opts ...grpc.CallOption) (*AnnouncementListResp, error) {
	out := new(AnnouncementListResp)
	err := c.cc.Invoke(ctx, Sms_GetAnnouncementList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, Sms_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServer is the server API for Sms service.
// All implementations must embed UnimplementedSmsServer
// for forward compatibility
type SmsServer interface {
	// 注册
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	// 登录
	Login(context.Context, *LoginReq) (*LoginResp, error)
	// 获取店铺
	GetShop(context.Context, *GetShopReq) (*GetShopResp, error)
	// 获取商品列表根据商品类型
	GetProductList(context.Context, *ProductListReq) (*ProductListResp, error)
	// 获取折扣商品列表根据商品类型
	GetLowProductList(context.Context, *LowProductListReq) (*LowProductListResp, error)
	// 获取详细商品信息
	GetDetailedProduct(context.Context, *DetailedProductReq) (*DetailedProductResp, error)
	// 计算总价
	GetTotalPrice(context.Context, *GetTotalPriceReq) (*GetTotalPriceResp, error)
	// 下单
	PostOrder(context.Context, *PostOrderReq) (*PostOrderResp, error)
	// 获取订单列表
	GetOrderList(context.Context, *GetOrderListReq) (*GetOrderListResp, error)
	// 查看订单详细信息
	GetDetailedOrder(context.Context, *DetailedOrderReq) (*DetailedOrderResp, error)
	// 取消订单
	OverOrder(context.Context, *ConfirmOrderReq) (*ConfirmOrderResp, error)
	// 取消申请
	CancellationOverOrder(context.Context, *ConfirmOrderReq) (*ConfirmOrderResp, error)
	// 确认订单
	ConfirmOrder(context.Context, *ConfirmOrderReq) (*ConfirmOrderResp, error)
	// 获取对应的店铺的公告列表
	GetAnnouncementList(context.Context, *AnnouncementListReq) (*AnnouncementListResp, error)
	// 获取个人信息
	GetUserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
	mustEmbedUnimplementedSmsServer()
}

// UnimplementedSmsServer must be embedded to have forward compatible implementations.
type UnimplementedSmsServer struct {
}

func (UnimplementedSmsServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedSmsServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSmsServer) GetShop(context.Context, *GetShopReq) (*GetShopResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShop not implemented")
}
func (UnimplementedSmsServer) GetProductList(context.Context, *ProductListReq) (*ProductListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductList not implemented")
}
func (UnimplementedSmsServer) GetLowProductList(context.Context, *LowProductListReq) (*LowProductListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLowProductList not implemented")
}
func (UnimplementedSmsServer) GetDetailedProduct(context.Context, *DetailedProductReq) (*DetailedProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailedProduct not implemented")
}
func (UnimplementedSmsServer) GetTotalPrice(context.Context, *GetTotalPriceReq) (*GetTotalPriceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalPrice not implemented")
}
func (UnimplementedSmsServer) PostOrder(context.Context, *PostOrderReq) (*PostOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostOrder not implemented")
}
func (UnimplementedSmsServer) GetOrderList(context.Context, *GetOrderListReq) (*GetOrderListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}
func (UnimplementedSmsServer) GetDetailedOrder(context.Context, *DetailedOrderReq) (*DetailedOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailedOrder not implemented")
}
func (UnimplementedSmsServer) OverOrder(context.Context, *ConfirmOrderReq) (*ConfirmOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OverOrder not implemented")
}
func (UnimplementedSmsServer) CancellationOverOrder(context.Context, *ConfirmOrderReq) (*ConfirmOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancellationOverOrder not implemented")
}
func (UnimplementedSmsServer) ConfirmOrder(context.Context, *ConfirmOrderReq) (*ConfirmOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmOrder not implemented")
}
func (UnimplementedSmsServer) GetAnnouncementList(context.Context, *AnnouncementListReq) (*AnnouncementListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncementList not implemented")
}
func (UnimplementedSmsServer) GetUserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedSmsServer) mustEmbedUnimplementedSmsServer() {}

// UnsafeSmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServer will
// result in compilation errors.
type UnsafeSmsServer interface {
	mustEmbedUnimplementedSmsServer()
}

func RegisterSmsServer(s grpc.ServiceRegistrar, srv SmsServer) {
	s.RegisterService(&Sms_ServiceDesc, srv)
}

func _Sms_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_GetShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetShop(ctx, req.(*GetShopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_GetProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetProductList(ctx, req.(*ProductListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetLowProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LowProductListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetLowProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_GetLowProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetLowProductList(ctx, req.(*LowProductListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetDetailedProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailedProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetDetailedProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_GetDetailedProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetDetailedProduct(ctx, req.(*DetailedProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetTotalPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalPriceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetTotalPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_GetTotalPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetTotalPrice(ctx, req.(*GetTotalPriceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_PostOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).PostOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_PostOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).PostOrder(ctx, req.(*PostOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_GetOrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetOrderList(ctx, req.(*GetOrderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetDetailedOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailedOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetDetailedOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_GetDetailedOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetDetailedOrder(ctx, req.(*DetailedOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_OverOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).OverOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_OverOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).OverOrder(ctx, req.(*ConfirmOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_CancellationOverOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).CancellationOverOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_CancellationOverOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).CancellationOverOrder(ctx, req.(*ConfirmOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_ConfirmOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).ConfirmOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_ConfirmOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).ConfirmOrder(ctx, req.(*ConfirmOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetAnnouncementList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnouncementListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetAnnouncementList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_GetAnnouncementList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetAnnouncementList(ctx, req.(*AnnouncementListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sms_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetUserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sms_ServiceDesc is the grpc.ServiceDesc for Sms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sms.Sms",
	HandlerType: (*SmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Sms_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Sms_Login_Handler,
		},
		{
			MethodName: "GetShop",
			Handler:    _Sms_GetShop_Handler,
		},
		{
			MethodName: "GetProductList",
			Handler:    _Sms_GetProductList_Handler,
		},
		{
			MethodName: "GetLowProductList",
			Handler:    _Sms_GetLowProductList_Handler,
		},
		{
			MethodName: "GetDetailedProduct",
			Handler:    _Sms_GetDetailedProduct_Handler,
		},
		{
			MethodName: "GetTotalPrice",
			Handler:    _Sms_GetTotalPrice_Handler,
		},
		{
			MethodName: "PostOrder",
			Handler:    _Sms_PostOrder_Handler,
		},
		{
			MethodName: "GetOrderList",
			Handler:    _Sms_GetOrderList_Handler,
		},
		{
			MethodName: "GetDetailedOrder",
			Handler:    _Sms_GetDetailedOrder_Handler,
		},
		{
			MethodName: "OverOrder",
			Handler:    _Sms_OverOrder_Handler,
		},
		{
			MethodName: "CancellationOverOrder",
			Handler:    _Sms_CancellationOverOrder_Handler,
		},
		{
			MethodName: "ConfirmOrder",
			Handler:    _Sms_ConfirmOrder_Handler,
		},
		{
			MethodName: "GetAnnouncementList",
			Handler:    _Sms_GetAnnouncementList_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Sms_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Sms.proto",
}
