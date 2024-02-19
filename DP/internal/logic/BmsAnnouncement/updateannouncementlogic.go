package BmsAnnouncement

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAnnouncementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAnnouncementLogic {
	return &UpdateAnnouncementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateAnnouncement 更新公告
func (l *UpdateAnnouncementLogic) UpdateAnnouncement(req *types.UpdateAnnouncementReq) (resp *types.BmsDataResp, err error) {
	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	updateAnnouncementReq := bmsclient.UpdateAnnouncementReq{
		StaffId:     StaffId,
		NoticeId:    req.NoticeId,
		NoticeTitle: req.NoticeTitle,
	}

	//调用RPC的服务
	updateAnnouncementResp, err := l.svcCtx.BmsRpcClient.UpdateAnnouncement(l.ctx, &updateAnnouncementReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(updateAnnouncementResp.Code),
		Msg:  updateAnnouncementResp.Msg,
		Data: updateAnnouncementResp.NoticeId,
	}, nil
}
