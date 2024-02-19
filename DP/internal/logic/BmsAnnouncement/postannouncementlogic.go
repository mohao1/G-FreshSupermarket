package BmsAnnouncement

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostAnnouncementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostAnnouncementLogic {
	return &PostAnnouncementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PostAnnouncement 发布公告
func (l *PostAnnouncementLogic) PostAnnouncement(req *types.PostAnnouncementReq) (resp *types.BmsDataResp, err error) {
	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	announcementReq := bmsclient.AnnouncementReq{
		StaffId:     StaffId,
		NoticeTitle: req.NoticeTitle,
	}

	//调用RPC的服务
	announcementResp, err := l.svcCtx.BmsRpcClient.PostAnnouncement(l.ctx, &announcementReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(announcementResp.Code),
		Msg:  announcementResp.Msg,
		Data: announcementResp.NoticeId,
	}, nil
}
