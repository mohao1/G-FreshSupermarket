package logic

import (
	"context"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PersonalInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersonalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersonalInfoLogic {
	return &PersonalInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PersonalInfo 获取个人信息
func (l *PersonalInfoLogic) PersonalInfo(in *bms.PersonalInfoReq) (*bms.PersonalInfoResp, error) {

	//查询身份信息
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//查询权限
	position, err := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
	if err != nil {
		return nil, err
	}

	//查询店铺信息
	shop, err := l.svcCtx.ShopModel.FindOne(l.ctx, staff.ShopId)
	if err != nil {
		return nil, err
	}

	return &bms.PersonalInfoResp{
		StaffId:      staff.StaffId,
		StaffName:    staff.StaffName,
		PositionName: position.PositionName,
		ShopId:       staff.ShopId,
		ShopName:     shop.ShopId,
	}, nil
}
