package logic

import (
	"context"
	"errors"
	"strconv"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSalesRecordsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopSalesRecordsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSalesRecordsListLogic {
	return &GetShopSalesRecordsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopSalesRecordsList 各个店铺商品总销售的数据列表
func (l *GetShopSalesRecordsListLogic) GetShopSalesRecordsList(in *ams.GetShopSalesRecordsListReq) (*ams.GetShopSalesRecordsListResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询信息
	salesRecordsList, err := l.svcCtx.OrderModel.SelectSalesRecordsListByShopId(l.ctx, in.ShopId)
	if err != nil {
		return nil, err
	}

	salesRecordsSumList := make([]*ams.SalesRecordsSumData, len(*salesRecordsList))

	for i, recordsData := range *salesRecordsList {
		if recordsData.ProductId[:3] == "LPR" {
			product, err := l.svcCtx.LowProductListModel.FindOne(l.ctx, recordsData.ProductId)
			salesRecordsSumList[i] = &ams.SalesRecordsSumData{
				ProductId:              recordsData.ProductId,
				ProductName:            product.ProductName,
				ProductPicture:         recordsData.ProductPrice,
				ProductSalesRecordsSum: strconv.Itoa(int(recordsData.ProductSalesRecordsSum)),
			}
			if err != nil {
				return nil, err
			}
		} else {
			product, err := l.svcCtx.ProductListModel.FindOne(l.ctx, recordsData.ProductId)
			salesRecordsSumList[i] = &ams.SalesRecordsSumData{
				ProductId:              recordsData.ProductId,
				ProductName:            product.ProductName,
				ProductPicture:         recordsData.ProductPrice,
				ProductSalesRecordsSum: strconv.Itoa(int(recordsData.ProductSalesRecordsSum)),
			}
			if err != nil {
				return nil, err
			}
		}

	}

	return &ams.GetShopSalesRecordsListResp{
		Code:                200,
		Msg:                 "获取成功",
		SalesRecordsSumList: salesRecordsSumList,
	}, nil
}
