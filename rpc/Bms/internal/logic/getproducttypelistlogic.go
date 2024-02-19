package logic

import (
	"context"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

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
func (l *GetProductTypeListLogic) GetProductTypeList(in *bms.GetProductTypeListReq) (*bms.GetProductTypeListResp, error) {

	typeList, err := l.svcCtx.ProductTypeModel.SelectTypeList(l.ctx)
	if err != nil {
		return nil, err
	}

	productTypeList := make([]*bms.ProductType, len(*typeList))

	for i, productType := range *typeList {
		productTypeList[i] = &bms.ProductType{
			ProductTypeId:   productType.ProductTypeId,
			ProductTypeName: productType.ProductTypeName,
			ProductUnit:     productType.ProductUnit,
		}
	}

	return &bms.GetProductTypeListResp{
		Code:        200,
		Msg:         "获取成功",
		ProductType: productTypeList,
	}, nil
}
