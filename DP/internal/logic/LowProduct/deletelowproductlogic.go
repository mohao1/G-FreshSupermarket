package LowProduct

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLowProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLowProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLowProductLogic {
	return &DeleteLowProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLowProductLogic) DeleteLowProduct(req *types.DeleteLowProductReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	deleteProduct := bmsclient.DeleteLowProductReq{
		StaffId:   StaffId,
		ProductId: req.ProductId,
	}
	//调用RPC的服务
	lowProductResp, err := l.svcCtx.BmsRpcClient.DeleteLowProduct(l.ctx, &deleteProduct)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(lowProductResp.Code),
		Msg:  lowProductResp.Msg,
		Data: lowProductResp.ProductId,
	}, nil
}
