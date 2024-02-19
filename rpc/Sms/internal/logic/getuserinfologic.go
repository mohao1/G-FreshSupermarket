package logic

import (
	"context"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserInfo 获取个人信息
func (l *GetUserInfoLogic) GetUserInfo(in *sms.UserInfoReq) (*sms.UserInfoResp, error) {

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	position, err := l.svcCtx.PositionModel.FindOne(l.ctx, user.PositionId)
	if err != nil {
		return nil, err
	}

	return &sms.UserInfoResp{
		UserId:           user.Id,
		UserName:         user.Name,
		PositionName:     position.PositionName,
		RegistrationTime: user.RegistrationTime.String(),
	}, nil
}
