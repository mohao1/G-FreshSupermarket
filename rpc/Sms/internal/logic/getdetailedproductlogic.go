package logic

import (
	"context"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailedProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDetailedProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailedProductLogic {
	return &GetDetailedProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetDetailedProduct 获取详细商品信息
func (l *GetDetailedProductLogic) GetDetailedProduct(in *sms.DetailedProductReq) (*sms.DetailedProductResp, error) {
	switch in.Type {
	case "low":
		findOne, err := l.svcCtx.LowProductList.FindOne(l.ctx, in.ProductId)
		if err != nil {
			return nil, err
		}
		productType, err := l.svcCtx.ProductType.FindOne(l.ctx, findOne.ProductTypeId)
		if err != nil {
			return nil, err
		}
		return &sms.DetailedProductResp{
			ProductId:      findOne.ProductId,
			ProductName:    findOne.ProductName,
			ProductTitle:   findOne.ProductTitle,
			ProductPicture: findOne.ProductPicture,
			ProductType:    productType.ProductTypeName,
			Price:          findOne.Price,
			ProductSize:    findOne.ProductSize,
			Producer:       findOne.Producer,
			ProductUnit:    productType.ProductUnit,
		}, nil
	case "list":

		findOne, err := l.svcCtx.ProductList.FindOne(l.ctx, in.ProductId)
		if err != nil {
			return nil, err
		}
		productType, err := l.svcCtx.ProductType.FindOne(l.ctx, findOne.ProductTypeId)
		if err != nil {
			return nil, err
		}

		return &sms.DetailedProductResp{
			ProductId:      findOne.ProductId,
			ProductName:    findOne.ProductName,
			ProductTitle:   findOne.ProductTitle,
			ProductPicture: findOne.ProductPicture,
			ProductType:    productType.ProductTypeName,
			Price:          findOne.Price,
			ProductSize:    findOne.ProductSize,
			Producer:       findOne.Producer,
			ProductUnit:    productType.ProductUnit,
		}, nil
	default:
		return nil, nil
	}
}
