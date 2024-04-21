package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSalesRecordsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopSalesRecordsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSalesRecordsListLogic {
	return &GetShopSalesRecordsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopSalesRecordsList 各个店铺商品总销售的数据列表
func (l *GetShopSalesRecordsListLogic) GetShopSalesRecordsList(req *types.GetShopSalesRecordsListReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopSalesRecordsListReq := amsclient.GetShopSalesRecordsListReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
	}

	//调用RPC的服务
	shopSalesRecordsList, err := l.svcCtx.AmsRpcClient.GetShopSalesRecordsList(l.ctx, &shopSalesRecordsListReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopSalesRecordsList.Code,
		Msg:  shopSalesRecordsList.Msg,
		Data: shopSalesRecordsList.SalesRecordsSumList,
	}, nil
}
