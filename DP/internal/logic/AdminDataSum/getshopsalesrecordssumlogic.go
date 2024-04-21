package AdminDataSum

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopSalesRecordsSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopSalesRecordsSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopSalesRecordsSumLogic {
	return &GetShopSalesRecordsSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopSalesRecordsSum 各个店铺总销售的数据
func (l *GetShopSalesRecordsSumLogic) GetShopSalesRecordsSum(req *types.GetShopSalesRecordsSumReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopSalesRecordsSumReq := amsclient.GetShopSalesRecordsSumReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
	}
	//调用RPC的服务
	shopSalesRecordsSum, err := l.svcCtx.AmsRpcClient.GetShopSalesRecordsSum(l.ctx, &shopSalesRecordsSumReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopSalesRecordsSum.Code,
		Msg:  shopSalesRecordsSum.Msg,
		Data: shopSalesRecordsSum.ShopSalesRecordsSum,
	}, nil
}
