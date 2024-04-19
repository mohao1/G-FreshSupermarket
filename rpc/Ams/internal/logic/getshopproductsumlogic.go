package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopProductSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopProductSumLogic {
	return &GetShopProductSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopProductSum 进行门店对应普通商品数量统计
func (l *GetShopProductSumLogic) GetShopProductSum(in *ams.GetShopProductSumReq) (*ams.GetShopProductSumResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询数据
	productSum, err := l.svcCtx.ProductListModel.SelectProductSumByShopId(l.ctx, in.ShopId)
	if err != nil {
		return nil, err
	}

	return &ams.GetShopProductSumResp{
		Code:       200,
		Msg:        "获取成功",
		ProductSum: productSum.CountNumber,
	}, nil
}
