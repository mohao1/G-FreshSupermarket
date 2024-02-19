package BmsSalesData

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSalesDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSalesDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSalesDataLogic {
	return &GetSalesDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetSalesData 查看销售数据
func (l *GetSalesDataLogic) GetSalesData(req *types.GetSalesDataReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	salesDataListReq := bmsclient.SalesDataListReq{
		StaffId: StaffId,
	}

	//调用RPC的服务
	salesDataListResp, err := l.svcCtx.BmsRpcClient.GetSalesData(l.ctx, &salesDataListReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(salesDataListResp.Code),
		Msg:  salesDataListResp.Msg,
		Data: salesDataListResp.SalesRecords,
	}, nil
}
