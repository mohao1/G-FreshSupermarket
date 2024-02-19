package Product

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateProductLogic {
	return &UpDateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpDateProduct 修改发布普通商品
func (l *UpDateProductLogic) UpDateProduct(req *types.UpDateProductReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	upDateProduct := bmsclient.UpDateProductReq{
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
	}

	//调用RPC的服务
	upDateProductResp, err := l.svcCtx.BmsRpcClient.UpDateProduct(l.ctx, &upDateProduct)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(upDateProductResp.Code),
		Msg:  upDateProductResp.Msg,
		Data: upDateProductResp.ProductId,
	}, nil
}
