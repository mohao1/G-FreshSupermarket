package Product

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReleaseProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReleaseProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseProductLogic {
	return &ReleaseProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReleaseProduct 设置发布普通商品
func (l *ReleaseProductLogic) ReleaseProduct(req *types.ReleaseProductReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	releaseProduct := bmsclient.ReleaseProductReq{
		StaffId:         StaffId,
		ProductName:     req.ProductName,
		ProductTitle:    req.ProductTitle,
		ProductTypeId:   req.ProductTypeId,
		ProductQuantity: req.ProductQuantity,
		ProductPicture:  req.ProductPicture,
		Price:           req.Price,
		ProductSize:     req.ProductSize,
		Producer:        req.Producer,
	}

	//调用RPC的服务
	releaseProductResp, err := l.svcCtx.BmsRpcClient.ReleaseProduct(l.ctx, &releaseProduct)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(releaseProductResp.Code),
		Msg:  releaseProductResp.Msg,
		Data: releaseProductResp.ProductId,
	}, nil
}
