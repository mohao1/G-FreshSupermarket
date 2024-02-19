package svc

import (
	"DP/rpc/Sms/internal/config"
	"DP/rpc/model"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	UserModel         model.UserModel
	ShopModel         model.ShopModel
	ProductList       model.ProductListModel
	LowProductList    model.LowProductListModel
	ProductType       model.ProductTypeModel
	OrderModel        model.OrderModel
	OrderNumberModel  model.OrderNumberModel
	RefundRecordModel model.RefundRecordModel
	PositionModel     model.PositionModel
	NoticeModel       model.NoticeModel
	//注册事务进入全局
	TransactCtx func(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:            c,
		UserModel:         model.NewUserModel(mysql),
		ShopModel:         model.NewShopModel(mysql),
		ProductList:       model.NewProductListModel(mysql),
		LowProductList:    model.NewLowProductListModel(mysql),
		ProductType:       model.NewProductTypeModel(mysql),
		OrderModel:        model.NewOrderModel(mysql),
		OrderNumberModel:  model.NewOrderNumberModel(mysql),
		RefundRecordModel: model.NewRefundRecordModel(mysql),
		PositionModel:     model.NewPositionModel(mysql),
		NoticeModel:       model.NewNoticeModel(mysql),
		//注册事务进入全局
		TransactCtx: mysql.TransactCtx,
	}
}
