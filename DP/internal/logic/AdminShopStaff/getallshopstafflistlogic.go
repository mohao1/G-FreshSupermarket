package AdminShopStaff

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllShopStaffListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllShopStaffListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllShopStaffListLogic {
	return &GetAllShopStaffListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetAllShopStaffList 查看全部店铺员工列表
func (l *GetAllShopStaffListLogic) GetAllShopStaffList(req *types.GetAllShopStaffListReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	shopStaffListReq := amsclient.GetAllShopStaffListReq{
		AdminId: AdminId,
		Limit:   req.Limit,
	}

	//调用RPC的服务
	shopStaffListResp, err := l.svcCtx.AmsRpcClient.GetAllShopStaffList(l.ctx, &shopStaffListReq)
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
