package logic

import (
	"context"
	"errors"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductDetailsLogic {
	return &GetProductDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductDetails 查看商品详细情况
func (l *GetProductDetailsLogic) GetProductDetails(in *bms.GetProductDetailsReq) (*bms.GetProductDetailsResp, error) {

	//查询店铺信息和个人的信息
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//进行权限判断
	position, err := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
	if err != nil {
		return nil, err
	}
	if position.PositionName != "经理" {
		return nil, errors.New("权限不足")
	}

	//查询对应信息
	product, err := l.svcCtx.ProductListModel.FindOne(l.ctx, in.ProductId)
	if err != nil {
		return nil, err
	}

	//店铺信息校验
	if product.ShopId != staff.ShopId {
		return nil, errors.New("店铺信息错误")
	}

	//查询商品类型信息
	productType, err := l.svcCtx.ProductTypeModel.FindOne(l.ctx, product.ProductTypeId)
	if err != nil {
		return nil, err
	}

	return &bms.GetProductDetailsResp{
		ProductId:       product.ProductId,
		ProductName:     product.ProductName,
		ProductTitle:    product.ProductTitle,
		ProductTypeName: productType.ProductTypeName,
		ProductUnit:     productType.ProductUnit,
		ProductQuantity: product.ProductQuantity,
		ProductPicture:  product.ProductPicture,
		Price:           product.Price,
		ProductSize:     product.ProductSize,
		Producer:        product.Producer,
	}, nil
}
