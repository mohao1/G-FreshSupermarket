package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductTypeListLogic {
	return &GetProductTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductTypeList 获取商品类型列表
func (l *GetProductTypeListLogic) GetProductTypeList(in *ams.GetProductTypeListReq) (*ams.GetProductTypeListResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询类型列表
	productTypes, err := l.svcCtx.ProductTypeModel.SelectTypeList(l.ctx)
	if err != nil {
		return nil, err
	}

	//准备数据信息
	positionList := make([]*ams.ProductType, len(*productTypes))

	for i, productType := range *productTypes {
		positionList[i] = &ams.ProductType{
			ProductTypeId:   productType.ProductTypeId,
			ProductTypeName: productType.ProductTypeName,
			ProductTypeUnit: productType.ProductUnit,
		}
	}

	return &ams.GetProductTypeListResp{
		Code:            200,
		Msg:             "获取成功",
		ProductTypeList: positionList,
	}, nil
}
