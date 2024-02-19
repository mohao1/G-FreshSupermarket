package logic

import (
	"DP/rpc/model"
	"context"
	"errors"
	"time"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAnnouncementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAnnouncementLogic {
	return &UpdateAnnouncementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateAnnouncement 更新公告
func (l *UpdateAnnouncementLogic) UpdateAnnouncement(in *bms.UpdateAnnouncementReq) (*bms.UpdateAnnouncementResp, error) {

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

	notice, err := l.svcCtx.NoticeModel.FindOne(l.ctx, in.NoticeId)
	if err != nil {
		return nil, err
	}

	if notice.ShopId != staff.ShopId {
		return nil, errors.New("店铺数据错误")
	}

	data := model.Notice{
		NoticeId:     in.NoticeId,
		NoticeTitle:  in.NoticeTitle,
		CreationTime: notice.CreationTime,
		UpdataTime:   time.Now(),
		ShopId:       notice.ShopId,
	}

	err = l.svcCtx.NoticeModel.Update(l.ctx, &data)
	if err != nil {
		return nil, err
	}

	return &bms.UpdateAnnouncementResp{
		Code:     200,
		Msg:      "修改成功",
		NoticeId: in.NoticeId,
	}, nil
}
