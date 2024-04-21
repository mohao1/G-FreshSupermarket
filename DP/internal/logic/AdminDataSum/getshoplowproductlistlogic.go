package AdminDataSum

import (
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopLowProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopLowProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopLowProductListLogic {
	return &GetShopLowProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopLowProductListLogic) GetShopLowProductList(req *types.GetShopLowProductListReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
