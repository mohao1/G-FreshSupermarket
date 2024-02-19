package logic

import (
	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLowProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductListLogic {
	return &GetLowProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLowProductList 获取折扣商品列表
func (l *GetLowProductListLogic) GetLowProductList(in *sms.LowProductListReq) (*sms.LowProductListResp, error) {
	switch in.ProductType {
	case "all":
		{
			limit, err := l.svcCtx.LowProductList.SelectAllLimit(l.ctx, in)
			if err != nil {
				return nil, err
			}

			lowProds := make([]*sms.LowProduct, len(*limit))

			for i, product := range *limit {
				lowProds[i] = &sms.LowProduct{
					ProductId:      product.ProductId,
					ProductName:    product.ProductName,
					ProductPicture: product.ProductPicture,
					ProductType:    product.ProductTypeName,
					Price:          product.Price,
					ProductSize:    product.ProductSize,
					Producer:       product.Producer,
					Quota:          product.Quota,
					StartTime:      product.StartTime.String(),
					EndTime:        product.EndTime.String(),
				}
			}

			return &sms.LowProductListResp{
				Code:        200,
				ProductList: lowProds,
				Msg:         "查询成功",
			}, nil

		}
	default:
		{
			//判断类型是否存在
			selectType, _ := l.svcCtx.ProductType.SelectType(l.ctx, in.ProductType)
			fmt.Println(selectType)
			if selectType == nil {
				return &sms.LowProductListResp{
					Code:        400,
					ProductList: nil,
					Msg:         "没有对应类型",
				}, nil
			}

			limit, err := l.svcCtx.LowProductList.SelectLowProductListLimit(l.ctx, in)
			if err != nil {
				return nil, err
			}

			lowProds := make([]*sms.LowProduct, len(*limit))

			for i, product := range *limit {
				lowProds[i] = &sms.LowProduct{
					ProductId:      product.ProductId,
					ProductName:    product.ProductName,
					ProductPicture: product.ProductPicture,
					ProductType:    product.ProductTypeName,
					Price:          product.Price,
					ProductSize:    product.ProductSize,
					Producer:       product.Producer,
					Quota:          product.Quota,
					StartTime:      product.StartTime.String(),
					EndTime:        product.EndTime.String(),
				}
			}

			return &sms.LowProductListResp{
				Code:        200,
				ProductList: lowProds,
				Msg:         "查询成功",
			}, nil
		}
	}
}
