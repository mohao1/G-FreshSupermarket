package logic

import (
	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSalesDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSalesDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSalesDataLogic {
	return &GetSalesDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetSalesData 查看销售数据
func (l *GetSalesDataLogic) GetSalesData(in *bms.SalesDataListReq) (*bms.SalesDataListResp, error) {

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

	list, err := l.svcCtx.SalesRecordsModel.SelectSalesRecordsListByShopId(l.ctx, staff.ShopId)
	if err != nil {

		return nil, err
	}

	salesRecordList := make([]*bms.SalesRecords, len(*list))

	for i, records := range *list {

		if records.ProductId[:3] == "LPR" {
			//处理特价商品
			product, err2 := l.svcCtx.LowProductListModel.FindOne(l.ctx, records.ProductId)
			if err2 != nil {
				return nil, err2
			}
			salesRecordList[i] = &bms.SalesRecords{
				SalesRecordsId: records.SalesRecordsId,
				ProductName:    product.ProductName,
				SalesQuantity:  records.SalesQuantity,
				TotalPrice:     records.TotalPrice,
			}
		} else {
			//处理普通商品
			product, err2 := l.svcCtx.ProductListModel.FindOne(l.ctx, records.ProductId)
			if err2 != nil {
				return nil, err2
			}
			salesRecordList[i] = &bms.SalesRecords{
				SalesRecordsId: records.SalesRecordsId,
				ProductName:    product.ProductName,
				SalesQuantity:  records.SalesQuantity,
				TotalPrice:     records.TotalPrice,
			}
		}

	}

	return &bms.SalesDataListResp{
		Code:         200,
		Msg:          "获取成功",
		SalesRecords: salesRecordList,
	}, nil
}
