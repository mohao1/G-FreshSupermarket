package AdminProductType

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductTypeLogic {
	return &DeleteProductTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductType 删除商品类型
func (l *DeleteProductTypeLogic) DeleteProductType(req *types.DeleteProductTypeReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	productTypeReq := amsclient.DeleteProductTypeReq{
		AdminId:       AdminId,
		ProductTypeId: req.ProductTypeId,
	}

	//调用RPC的服务
	productTypeResp, err := l.svcCtx.AmsRpcClient.DeleteProductType(l.ctx, &productTypeReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: productTypeResp.Code,
		Msg:  productTypeResp.Msg,
		Data: productTypeResp.ProductTypeId,
	}, nil

}
