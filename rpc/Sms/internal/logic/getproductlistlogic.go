package logic

import (
	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"
	"context"

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

// GetProductList 获取商品列表
func (l *GetProductListLogic) GetProductList(in *sms.ProductListReq) (*sms.ProductListResp, error) {
	switch in.ProductType {
	case "all":
		{
			limit, err := l.svcCtx.ProductList.SelectAllLimit(l.ctx, in)
			if err != nil {
				return nil, err
			}
			prods := make([]*sms.Product, len(*limit))
			for i, prod := range *limit {
				prods[i] = &sms.Product{
					ProductId:      prod.ProductId,
					ProductName:    prod.ProductName,
					ProductPicture: prod.ProductPicture,
					ProductType:    prod.ProductTypeName,
					Price:          prod.Price,
					ProductSize:    prod.ProductSize,
					Producer:       prod.Producer,
				}
			}
			return &sms.ProductListResp{
				Code:        200,
				ProductList: prods,
				Msg:         "查询成功",
			}, nil
		}
	default:
		{
			selectType, _ := l.svcCtx.ProductType.SelectType(l.ctx, in.ProductType)
			if selectType == nil {
				return &sms.ProductListResp{
					Code:        400,
					ProductList: nil,
					Msg:         "没有对应类型",
				}, nil
			}
			limit, err := l.svcCtx.ProductList.SelectProductListLimit(l.ctx, in)
			if err != nil {
				return nil, err
			}
			prods := make([]*sms.Product, len(*limit))
			for i, prod := range *limit {
				prods[i] = &sms.Product{
					ProductId:      prod.ProductId,
					ProductName:    prod.ProductName,
					ProductPicture: prod.ProductPicture,
					ProductType:    prod.ProductTypeName,
					Price:          prod.Price,
					ProductSize:    prod.ProductSize,
					Producer:       prod.Producer,
				}
			}
			return &sms.ProductListResp{
				Code:        200,
				ProductList: prods,
				Msg:         "查询成功",
			}, nil
		}

	}
}
