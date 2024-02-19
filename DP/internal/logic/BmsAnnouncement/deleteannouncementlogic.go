package BmsAnnouncement

import (
	"DP/rpc/Bms/bmsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAnnouncementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAnnouncementLogic {
	return &DeleteAnnouncementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteAnnouncement 删除公告
func (l *DeleteAnnouncementLogic) DeleteAnnouncement(req *types.DeleteAnnouncementReq) (resp *types.BmsDataResp, err error) {
	//获取JWT中的UserId
	StaffId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据信息
	deleteAnnouncementReq := bmsclient.DeleteAnnouncementReq{
		StaffId:  StaffId,
		NoticeId: req.NoticeId,
	}

	//调用RPC的服务
	deleteAnnouncementResp, err := l.svcCtx.BmsRpcClient.DeleteAnnouncement(l.ctx, &deleteAnnouncementReq)
	if err != nil {
		return &types.BmsDataResp{
			Code: 400,
			Msg:  "数据获取错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.BmsDataResp{
		Code: int(deleteAnnouncementResp.Code),
		Msg:  deleteAnnouncementResp.Msg,
		Data: deleteAnnouncementResp.NoticeId,
	}, nil
}
