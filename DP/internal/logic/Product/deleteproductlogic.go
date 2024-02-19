package Product

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProduct 删除发布普通商品
func (l *DeleteProductLogic) DeleteProduct(req *types.DeleteProductReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	deleteProduct := bmsclient.DeleteProductReq{
		StaffId:   StaffId,
		ProductId: req.ProductId,
	}

	//调用RPC的服务
	deleteProductResp, err := l.svcCtx.BmsRpcClient.DeleteProduct(l.ctx, &deleteProduct)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(deleteProductResp.Code),
		Msg:  deleteProductResp.Msg,
		Data: deleteProductResp.ProductId,
	}, nil
}
