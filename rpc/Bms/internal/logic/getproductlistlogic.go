package logic

import (
	"context"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductList 查看普通商品列表
func (l *GetProductListLogic) GetProductList(in *bms.GetProductListReq) (*bms.GetProductListResp, error) {

	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	products, err := l.svcCtx.ProductListModel.SelectProductListByShopId(l.ctx, nil, staff.ShopId, int(in.Limit))
	if err != nil {
		return nil, err
	}

	productList := make([]*bms.Product, len(*products))

	for i, list := range *products {
		productType, err2 := l.svcCtx.ProductTypeModel.FindOne(l.ctx, list.ProductTypeId)
		if err2 != nil {
			return nil, err2
		}

		productList[i] = &bms.Product{
			ProductId:       list.ProductId,
			ProductName:     list.ProductName,
			ProductTypeName: productType.ProductTypeName,
			ProductQuantity: list.ProductQuantity,
			ProductPicture:  list.ProductPicture,
			Producer:        list.Producer,
		}
	}

	return &bms.GetProductListResp{
		Code:    200,
		Msg:     "获取成功",
		Product: productList,
	}, nil
}
