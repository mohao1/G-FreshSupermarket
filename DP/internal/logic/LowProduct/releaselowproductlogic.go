package LowProduct

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReleaseLowProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReleaseLowProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseLowProductLogic {
	return &ReleaseLowProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReleaseLowProduct 设置发布普通商品
func (l *ReleaseLowProductLogic) ReleaseLowProduct(req *types.ReleaseLowProductReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	lowProductReq := bmsclient.ReleaseLowProductReq{
		StaffId:         StaffId,
		ProductName:     req.ProductName,
		ProductTitle:    req.ProductTitle,
		ProductTypeId:   req.ProductTypeId,
		ProductQuantity: req.ProductQuantity,
		ProductPicture:  req.ProductPicture,
		Price:           req.Price,
		ProductSize:     req.ProductSize,
		Producer:        req.Producer,
		Quota:           req.Quota,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
	}

	//调用RPC的服务
	releaseLowProductResp, err := l.svcCtx.BmsRpcClient.ReleaseLowProduct(l.ctx, &lowProductReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(releaseLowProductResp.Code),
		Msg:  releaseLowProductResp.Msg,
		Data: releaseLowProductResp.ProductId,
	}, nil
}
