package ProductType

import (
	"DP/rpc/Bms/bmsclient"
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

func (l *GetProductTypeListLogic) GetProductTypeList(req *types.GetProductTypeListReq) (resp *types.BmsDataResp, err error) {

	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	productTypeList := bmsclient.GetProductTypeListReq{
		StaffId: StaffId,
	}

	//调用RPC的服务
	typeListResp, err := l.svcCtx.BmsRpcClient.GetProductTypeList(l.ctx, &productTypeList)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "获取信息错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(typeListResp.Code),
		Msg:  typeListResp.Msg,
		Data: typeListResp.ProductType,
	}, nil
}
