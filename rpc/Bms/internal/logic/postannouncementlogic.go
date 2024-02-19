package logic

import (
	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"
	"DP/rpc/model"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostAnnouncementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostAnnouncementLogic {
	return &PostAnnouncementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PostAnnouncement 发布公告
func (l *PostAnnouncementLogic) PostAnnouncement(in *bms.AnnouncementReq) (*bms.AnnouncementResp, error) {

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

	AnnouncementId := uuid.New().String()

	data := model.Notice{
		NoticeId:     AnnouncementId,
		NoticeTitle:  in.NoticeTitle,
		CreationTime: time.Now(),
		UpdataTime:   time.Now(),
		ShopId:       staff.ShopId,
	}

	_, err = l.svcCtx.NoticeModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}

	return &bms.AnnouncementResp{
		Code:     200,
		Msg:      "发布成功",
		NoticeId: AnnouncementId,
	}, nil
}
