package AdminShopStaff

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopAllStaffSumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShopAllStaffSumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopAllStaffSumListLogic {
	return &GetShopAllStaffSumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetShopAllStaffSumList 店铺查询人员列表
func (l *GetShopAllStaffSumListLogic) GetShopAllStaffSumList(req *types.GetShopAllStaffSumListReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	staffSumListReq := amsclient.GetShopAllStaffSumListReq{
		AdminId: AdminId,
	}

	//调用RPC的服务
	staffSumListResp, err := l.svcCtx.AmsRpcClient.GetShopAllStaffSumList(l.ctx, &staffSumListReq)

	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: staffSumListResp.Code,
		Msg:  staffSumListResp.Msg,
		Data: staffSumListResp.StaffSumList,
	}, nil
}
