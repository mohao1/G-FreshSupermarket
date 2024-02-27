package AdminProductType

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductTypeListLogic {
	return &GetProductTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetProductTypeList 获取商品类型列表
func (l *GetProductTypeListLogic) GetProductTypeList(req *types.GetProductTypeListAdminReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	productTypeListReq := amsclient.GetProductTypeListReq{
		AdminId: AdminId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	productTypeListResp, err := l.svcCtx.AmsRpcClient.GetProductTypeList(l.ctx, &productTypeListReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: productTypeListResp.Code,
		Msg:  productTypeListResp.Msg,
		Data: productTypeListResp.ProductTypeList,
	}, nil
}
