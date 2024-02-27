package AdminProductType

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateProductTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateProductTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateProductTypeLogic {
	return &UpDateProductTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpDateProductType 修改商品类型
func (l *UpDateProductTypeLogic) UpDateProductType(req *types.UpDateProductTypeReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	productTypeReq := amsclient.UpDateProductTypeReq{
		AdminId:         AdminId,
		ProductTypeId:   req.ProductTypeId,
		ProductTypeName: req.ProductTypeName,
		ProductTypeUnit: req.ProductTypeUnit,
	}

	//调用RPC的服务
	productTypeResp, err := l.svcCtx.AmsRpcClient.UpDateProductType(l.ctx, &productTypeReq)
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
