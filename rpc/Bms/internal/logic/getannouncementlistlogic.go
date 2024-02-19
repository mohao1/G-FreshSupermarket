package logic

import (
	"context"
	"errors"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

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

// GetAnnouncementList 获取公告列表
func (l *GetAnnouncementListLogic) GetAnnouncementList(in *bms.AnnouncementListReq) (*bms.AnnouncementListResp, error) {

	//查询店铺信息和个人的信息
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//进行权限判断
	position, err := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
	if err != nil {
		return nil, err
	}
	if position.PositionName != "经理" {
		return nil, errors.New("权限不足")
	}

	list, err := l.svcCtx.NoticeModel.SelectNoticeByShopId(l.ctx, staff.ShopId)
	if err != nil {
		return nil, err
	}

	//创建数组
	AnnouncementList := make([]*bms.Announcement, len(*list))

	//填充数组
	for i, notice := range *list {
		AnnouncementList[i] = &bms.Announcement{
			NoticeId:     notice.NoticeId,
			NoticeTitle:  notice.NoticeTitle,
			CreationTime: notice.CreationTime.String(),
			UpdataTime:   notice.UpdataTime.String(),
		}
	}

	return &bms.AnnouncementListResp{
		Code:             200,
		Msg:              "获取成功",
		AnnouncementList: AnnouncementList,
	}, nil
}
