package svc

import (
	"DP/rpc/Ams/internal/config"
	"DP/rpc/model"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	AdminModel       model.AdminModel
	ShopModel        model.ShopModel
	StaffModel       model.StaffModel
	PositionModel    model.PositionModel
	ProductTypeModel model.ProductTypeModel
	UserModel        model.UserModel
	//注册事务进入全局
	TransactCtx func(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:           c,
		AdminModel:       model.NewAdminModel(mysql),
		ShopModel:        model.NewShopModel(mysql),
		StaffModel:       model.NewStaffModel(mysql),
		PositionModel:    model.NewPositionModel(mysql),
		ProductTypeModel: model.NewProductTypeModel(mysql),
		UserModel:        model.NewUserModel(mysql),
		//注册事务进入全局
		TransactCtx: mysql.TransactCtx,
	}
}
