package AdminDataSum

import (
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopProductListLogic {
	return &GetShopProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShopProductListLogic) GetShopProductList(req *types.GetShopProductListReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
