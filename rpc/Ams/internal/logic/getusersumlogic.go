package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSumLogic {
	return &GetUserSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserSum 用户人数
func (l *GetUserSumLogic) GetUserSum(in *ams.GetUserSumReq) (*ams.GetUserSumResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询信息
	userCount, err := l.svcCtx.UserModel.SelectUserListCount(l.ctx)
	if err != nil {
		return nil, err
	}

	return &ams.GetUserSumResp{
		Code:    200,
		Msg:     "获取成功",
		UserSum: userCount.UserCount,
	}, nil
}
