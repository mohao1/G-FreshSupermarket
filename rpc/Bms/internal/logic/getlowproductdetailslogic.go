package logic

import (
	"context"
	"errors"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLowProductDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductDetailsLogic {
	return &GetLowProductDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLowProductDetails 查看商品详细情况
func (l *GetLowProductDetailsLogic) GetLowProductDetails(in *bms.GetLowProductDetailsReq) (*bms.GetLowProductDetailsResp, error) {
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
	lowProduct, err := l.svcCtx.LowProductListModel.FindOne(l.ctx, in.ProductId)
	if err != nil {
		return nil, err
	}

	//店铺信息校验
	if lowProduct.ShopId != staff.ShopId {
		return nil, errors.New("店铺信息错误")
	}

	//查询商品类型信息
	productType, err := l.svcCtx.ProductTypeModel.FindOne(l.ctx, lowProduct.ProductTypeId)
	if err != nil {
		return nil, err
	}

	return &bms.GetLowProductDetailsResp{
		ProductId:       lowProduct.ProductId,
		ProductName:     lowProduct.ProductName,
		ProductTitle:    lowProduct.ProductTitle,
		ProductTypeName: productType.ProductTypeName,
		ProductUnit:     productType.ProductUnit,
		ProductQuantity: lowProduct.ProductQuantity,
		ProductPicture:  lowProduct.ProductPicture,
		Price:           lowProduct.Price,
		ProductSize:     lowProduct.ProductSize,
		Producer:        lowProduct.Producer,
		Quota:           lowProduct.Quota,
		StartTime:       lowProduct.StartTime.String(),
		EndTime:         lowProduct.EndTime.String(),
	}, nil

}
