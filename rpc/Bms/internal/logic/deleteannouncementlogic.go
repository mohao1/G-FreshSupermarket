package logic

import (
	"context"
	"errors"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAnnouncementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAnnouncementLogic {
	return &DeleteAnnouncementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteAnnouncement 删除公告
func (l *DeleteAnnouncementLogic) DeleteAnnouncement(in *bms.DeleteAnnouncementReq) (*bms.DeleteAnnouncementResp, error) {

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

	err = l.svcCtx.NoticeModel.Delete(l.ctx, in.NoticeId)
	if err != nil {
		return nil, err
	}

	return &bms.DeleteAnnouncementResp{
		Code:     200,
		Msg:      "删除成功",
		NoticeId: in.NoticeId,
	}, nil
}
