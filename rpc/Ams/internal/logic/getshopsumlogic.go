package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSumLogic {
	return &GetShopSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopSum 店铺数量
func (l *GetShopSumLogic) GetShopSum(in *ams.GetShopSumReq) (*ams.GetShopSumResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询信息
	shopListCount, err := l.svcCtx.ShopModel.SelectShopListCount(l.ctx)
	if err != nil {
		return nil, err
	}

	return &ams.GetShopSumResp{
		Code:    200,
		Msg:     "获取成功",
		ShopSum: shopListCount.ShopCount,
	}, nil
}
