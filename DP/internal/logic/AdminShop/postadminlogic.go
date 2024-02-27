package AdminShop

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostAdminLogic {
	return &PostAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostAdminLogic) PostAdmin(req *types.PostAdminReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	adminReq := amsclient.PostAdminReq{
		AdminId:   AdminId,
		StaffId:   req.StaffId,
		StaffName: req.StaffName,
		Password:  req.Password,
	}

	//调用RPC的服务
	adminResp, err := l.svcCtx.AmsRpcClient.PostAdmin(l.ctx, &adminReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: adminResp.Code,
		Msg:  adminResp.Msg,
		Data: adminResp.StaffId,
	}, nil
}
