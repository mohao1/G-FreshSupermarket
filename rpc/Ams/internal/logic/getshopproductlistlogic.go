package logic

import (
	"context"
	"errors"
	"strconv"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopProductListLogic {
	return &GetShopProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopProductList 获取门店对应普通商品列表
func (l *GetShopProductListLogic) GetShopProductList(in *ams.GetShopProductListReq) (*ams.GetShopProductListResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询数据
	products, err := l.svcCtx.ProductListModel.SelectProductListByShopId(l.ctx, nil, in.ShopId, int(in.Limit))
	if err != nil {
		return nil, err
	}

	productList := make([]*ams.PositionDataSc, len(*products))

	for i, product := range *products {

		shop, err2 := l.svcCtx.ShopModel.FindOne(l.ctx, product.ShopId)
		if err2 != nil {
			return nil, err2
		}
		productType, err2 := l.svcCtx.ProductTypeModel.FindOne(l.ctx, product.ProductTypeId)
		if err2 != nil {
			return nil, err2
		}

		productList[i] = &ams.PositionDataSc{
			ProductId:      product.ProductId,
			ProductName:    product.ProductName,
			Price:          product.Price,
			ProductSize:    strconv.Itoa(int(product.ProductSize)),
			ShopId:         product.ShopId,
			ShopName:       shop.ShopName,
			ProductPicture: product.ProductPicture,
			ProductType:    productType.ProductTypeName,
			CreationTime:   product.CreationTime.String(),
		}
	}

	return &ams.GetShopProductListResp{
		Code:        200,
		Msg:         "获取成功",
		ProductList: productList,
	}, nil
}
