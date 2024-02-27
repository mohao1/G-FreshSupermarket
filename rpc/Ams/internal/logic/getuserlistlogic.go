package logic

import (
	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserList 用户信息列表
func (l *GetUserListLogic) GetUserList(in *ams.GetUserListReq) (*ams.GetUserListResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询用户列表
	users, err := l.svcCtx.UserModel.TransactSelectUserList(l.ctx, nil, in.Limit)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, errors.New("查询数据错误")
	}

	//准备数据信息
	userList := make([]*ams.UserData, len(*users))

	for i, user := range *users {

		position, err2 := l.svcCtx.PositionModel.FindOne(l.ctx, user.PositionId)
		if err2 != nil {
			return nil, err2
		}

		userList[i] = &ams.UserData{
			UserId:           user.Id,
			UserName:         user.Name,
			Phone:            user.Phone,
			PositionId:       user.PositionId,
			PositionName:     position.PositionName,
			RegistrationTime: user.RegistrationTime.String(),
		}
	}

	return &ams.GetUserListResp{
		Code:         200,
		Msg:          "获取成功",
		UserDataList: userList,
	}, nil
}
