package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductTypeLogic {
	return &DeleteProductTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductType 删除商品类型
func (l *DeleteProductTypeLogic) DeleteProductType(in *ams.DeleteProductTypeReq) (*ams.DeleteProductTypeResp, error) {

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

		//删除类型信息数据
		err2 = l.svcCtx.ProductTypeModel.TransactDeleteProductType(ctx, session, in.ProductTypeId)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &ams.DeleteProductTypeResp{
		Code:          200,
		Msg:           "删除成功",
		ProductTypeId: in.ProductTypeId,
	}, nil
}
