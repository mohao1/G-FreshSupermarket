package logic

import (
	"DP/rpc/Utile"
	"DP/rpc/model"
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostProductTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostProductTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostProductTypeLogic {
	return &PostProductTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PostProductType 设置商品类型
func (l *PostProductTypeLogic) PostProductType(in *ams.PostProductTypeReq) (*ams.PostProductTypeResp, error) {
	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//生成随机Id
	productTypeId := "PT" + Utile.GetUUIDBy8()

	//准备插入数据
	data := model.ProductType{
		ProductTypeId:   productTypeId,
		ProductTypeName: in.ProductTypeName,
		ProductUnit:     in.ProductTypeUnit,
	}

	//插入数据信息
	_, err = l.svcCtx.ProductTypeModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}

	return &ams.PostProductTypeResp{
		Code:          200,
		Msg:           "设置成功",
		ProductTypeId: productTypeId,
	}, nil
}
