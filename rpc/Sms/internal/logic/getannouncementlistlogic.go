package logic

import (
	"context"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnnouncementListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnnouncementListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnnouncementListLogic {
	return &GetAnnouncementListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetAnnouncementList 获取对应的店铺的公告列表
func (l *GetAnnouncementListLogic) GetAnnouncementList(in *sms.AnnouncementListReq) (*sms.AnnouncementListResp, error) {
	NoticeList, err := l.svcCtx.NoticeModel.SelectNoticeByShopId(l.ctx, in.ShopId)
	if err != nil {
		return nil, err
	}
	announcementList := make([]*sms.Announcement, len(*NoticeList))

	for i, notice := range *NoticeList {
		announcementList[i] = &sms.Announcement{
			NoticeId:    notice.NoticeId,
			NoticeTitle: notice.NoticeTitle,
		}
	}

	if *NoticeList == nil {
		return &sms.AnnouncementListResp{
			Code:             200,
			Msg:              "没有规则",
			AnnouncementList: announcementList,
		}, nil
	}

	return &sms.AnnouncementListResp{
		Code:             200,
		Msg:              "获取成功",
		AnnouncementList: announcementList,
	}, nil
}
