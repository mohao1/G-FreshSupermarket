package LowProduct

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLowProductDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductDetailsLogic {
	return &GetLowProductDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetLowProductDetails 查看商品详细情况
func (l *GetLowProductDetailsLogic) GetLowProductDetails(req *types.GetLowProductDetailsReq) (resp *types.BmsDataResp, err error) {
	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	productDetailsReq := bmsclient.GetLowProductDetailsReq{
		StaffId:   StaffId,
		ProductId: req.ProductId,
	}

	//调用RPC的服务
	productDetailsResp, err := l.svcCtx.BmsRpcClient.GetLowProductDetails(l.ctx, &productDetailsReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: 200,
		Msg:  "获取成功",
		Data: productDetailsResp,
	}, nil
}
