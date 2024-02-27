package logic

import (
	"DP/rpc/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateProductTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpDateProductTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateProductTypeLogic {
	return &UpDateProductTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpDateProductType 修改商品类型
func (l *UpDateProductTypeLogic) UpDateProductType(in *ams.UpDateProductTypeReq) (*ams.UpDateProductTypeResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//查询类型信息数据
		productType, err2 := l.svcCtx.ProductTypeModel.TransactSelectProductType(ctx, session, in.ProductTypeId)
		if err2 != nil {
			return err2
		}
		if productType == nil {
			return errors.New("设置信息错误")
		}

		//准备数据
		data := model.ProductType{
			ProductTypeId:   in.ProductTypeId,
			ProductTypeName: in.ProductTypeName,
			ProductUnit:     in.ProductTypeUnit,
		}

		//修改类型数据信息
		err2 = l.svcCtx.ProductTypeModel.TransactUpDateProductType(ctx, session, &data)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &ams.UpDateProductTypeResp{
		Code:          200,
		Msg:           "修改成功",
		ProductTypeId: in.ProductTypeId,
	}, nil
}
