package AdminShopStaff

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopStaffListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopStaffListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopStaffListLogic {
	return &GetShopStaffListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopStaffList 各个门店人数统计
func (l *GetShopStaffListLogic) GetShopStaffList(req *types.GetShopStaffListReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopStaffListReq := amsclient.GetShopStaffListReq{
		AdminId: AdminId,
		ShopId:  req.ShopId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	shopStaffListResp, err := l.svcCtx.AmsRpcClient.GetShopStaffList(l.ctx, &shopStaffListReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: shopStaffListResp.Code,
		Msg:  shopStaffListResp.Msg,
		Data: shopStaffListResp.StaffList,
	}, nil
}
