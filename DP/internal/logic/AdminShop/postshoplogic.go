package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostShopLogic {
	return &PostShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostShopLogic) PostShop(req *types.PostShopListReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	postShopReq := amsclient.PostShopReq{
		AdminId:     AdminId,
		ShopId:      req.ShopId,
		ShopName:    req.ShopName,
		ShopAddress: req.ShopAddress,
		ShopCity:    req.ShopCity,
	}

	//调用RPC的服务
	postShopResp, err := l.svcCtx.AmsRpcClient.PostShop(l.ctx, &postShopReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: postShopResp.Code,
		Msg:  postShopResp.Msg,
		Data: postShopResp.ShopId,
	}, nil
}
