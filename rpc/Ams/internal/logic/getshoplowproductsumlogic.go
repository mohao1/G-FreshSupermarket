package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopLowProductSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopLowProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopLowProductSumLogic {
	return &GetShopLowProductSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopLowProductSum 进行门店对应折扣商品数量统计
func (l *GetShopLowProductSumLogic) GetShopLowProductSum(in *ams.GetShopLowProductSumReq) (*ams.GetShopLowProductSumResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询数据
	LowProductSum, err := l.svcCtx.LowProductListModel.SelectLowProductSumByShopId(l.ctx, in.ShopId)
	if err != nil {
		return nil, err
	}

	return &ams.GetShopLowProductSumResp{
		Code:       200,
		Msg:        "获取成功",
		ProductSum: LowProductSum.CountNumber,
	}, nil
}
