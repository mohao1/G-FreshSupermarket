package LowProduct

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateLowProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateLowProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateLowProductLogic {
	return &UpDateLowProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpDateLowProduct 修改发布普通商品
func (l *UpDateLowProductLogic) UpDateLowProduct(req *types.UpDateLowProductReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	lowProductReq := bmsclient.UpDateLowProductReq{
		StaffId:         StaffId,
		ProductId:       req.ProductId,
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
	upDateLowProductResp, err := l.svcCtx.BmsRpcClient.UpDateLowProduct(l.ctx, &lowProductReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(upDateLowProductResp.Code),
		Msg:  upDateLowProductResp.Msg,
		Data: upDateLowProductResp.ProductId,
	}, nil
}
