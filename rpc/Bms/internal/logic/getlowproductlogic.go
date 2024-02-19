package logic

import (
	"context"
	"fmt"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLowProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductLogic {
	return &GetLowProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLowProduct 查看折扣商品列表
func (l *GetLowProductLogic) GetLowProduct(in *bms.GetLowProductReq) (*bms.GetLowProductResp, error) {

	//信息查询
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	lowProducts, err := l.svcCtx.LowProductListModel.SelectLowProductListByShopId(l.ctx, nil, staff.ShopId, in.Limit)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	//创建信息列表
	lowProductList := make([]*bms.LowProduct, len(*lowProducts))

	//填充信息列表
	for i, lowProduct := range *lowProducts {

		//查询类型编号对应类型名称
		productType, err2 := l.svcCtx.ProductTypeModel.FindOne(l.ctx, lowProduct.ProductTypeId)
		if err2 != nil {
			return nil, err2
		}

		lowProductList[i] = &bms.LowProduct{
			ProductId:       lowProduct.ProductId,
			ProductName:     lowProduct.ProductName,
			ProductTypeName: productType.ProductTypeName,
			ProductQuantity: lowProduct.ProductQuantity,
			ProductPicture:  lowProduct.ProductPicture,
			Producer:        lowProduct.Producer,
		}
	}

	return &bms.GetLowProductResp{
		Code:       200,
		Msg:        "获取成功",
		LowProduct: lowProductList,
	}, nil
}
