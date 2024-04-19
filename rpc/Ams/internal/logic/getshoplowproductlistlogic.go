package logic

import (
	"context"
	"errors"
	"strconv"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopLowProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopLowProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopLowProductListLogic {
	return &GetShopLowProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopLowProductList 获取门店对应折扣商品列表
func (l *GetShopLowProductListLogic) GetShopLowProductList(in *ams.GetShopLowProductListReq) (*ams.GetShopLowProductListResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询数据
	lowProducts, err := l.svcCtx.LowProductListModel.SelectLowProductListByShopId(l.ctx, nil, in.ShopId, in.Limit)
	if err != nil {
		return nil, err
	}

	lowProductList := make([]*ams.PositionDataSc, len(*lowProducts))

	for i, lowProduct := range *lowProducts {

		shop, err2 := l.svcCtx.ShopModel.FindOne(l.ctx, lowProduct.ShopId)
		if err2 != nil {
			return nil, err2
		}
		productType, err2 := l.svcCtx.ProductTypeModel.FindOne(l.ctx, lowProduct.ProductTypeId)
		if err2 != nil {
			return nil, err2
		}

		lowProductList[i] = &ams.PositionDataSc{
			ProductId:      lowProduct.ProductId,
			ProductName:    lowProduct.ProductName,
			Price:          lowProduct.Price,
			ProductSize:    strconv.Itoa(int(lowProduct.ProductSize)),
			ShopId:         lowProduct.ShopId,
			ShopName:       shop.ShopName,
			ProductPicture: lowProduct.ProductPicture,
			ProductType:    productType.ProductTypeName,
			CreationTime:   lowProduct.CreationTime.String(),
		}
	}

	return &ams.GetShopLowProductListResp{
		Code:        200,
		Msg:         "获取成功",
		ProductList: lowProductList,
	}, nil
}
