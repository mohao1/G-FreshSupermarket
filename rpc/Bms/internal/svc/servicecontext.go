package svc

import (
	"DP/rpc/Bms/internal/config"
	"DP/rpc/model"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	StaffModel          model.StaffModel
	PositionModel       model.PositionModel
	ShopModel           model.ShopModel
	OrderNumberModel    model.OrderNumberModel
	UserModel           model.UserModel
	OrderModel          model.OrderModel
	ProductTypeModel    model.ProductTypeModel
	ProductListModel    model.ProductListModel
	LowProductListModel model.LowProductListModel
	RefundRecordModel   model.RefundRecordModel
	NoticeModel         model.NoticeModel
	SalesRecordsModel   model.SalesRecordsModel
	//注册事务进入全局
	TransactCtx func(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:              c,
		StaffModel:          model.NewStaffModel(mysql),
		PositionModel:       model.NewPositionModel(mysql),
		ShopModel:           model.NewShopModel(mysql),
		OrderNumberModel:    model.NewOrderNumberModel(mysql),
		UserModel:           model.NewUserModel(mysql),
		OrderModel:          model.NewOrderModel(mysql),
		ProductTypeModel:    model.NewProductTypeModel(mysql),
		ProductListModel:    model.NewProductListModel(mysql),
		LowProductListModel: model.NewLowProductListModel(mysql),
		RefundRecordModel:   model.NewRefundRecordModel(mysql),
		NoticeModel:         model.NewNoticeModel(mysql),
		SalesRecordsModel:   model.NewSalesRecordsModel(mysql),
		//注册事务进入全局
		TransactCtx: mysql.TransactCtx,
	}
}
