package BmsAnnouncement

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnnouncementListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAnnouncementListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnnouncementListLogic {
	return &GetAnnouncementListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetAnnouncementList 取公告列表
func (l *GetAnnouncementListLogic) GetAnnouncementList(req *types.GetAnnouncementListReq) (resp *types.BmsDataResp, err error) {
	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	announcementListReq := bmsclient.AnnouncementListReq{
		StaffId: StaffId,
	}

	//调用RPC的服务
	announcementListResp, err := l.svcCtx.BmsRpcClient.GetAnnouncementList(l.ctx, &announcementListReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(announcementListResp.Code),
		Msg:  announcementListResp.Msg,
		Data: announcementListResp.AnnouncementList,
	}, nil
}
