package logic

import (
	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNewUserSumToDayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNewUserSumToDayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewUserSumToDayLogic {
	return &GetNewUserSumToDayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetNewUserSumToDay 今日新增用户数量
func (l *GetNewUserSumToDayLogic) GetNewUserSumToDay(in *ams.GetNewUserSumToDayReq) (*ams.GetNewUserSumToDayResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询信息
	userNumber, err := l.svcCtx.UserModel.SelectUserNumberTheDay(l.ctx)
	if err != nil {
		return nil, err
	}

	return &ams.GetNewUserSumToDayResp{
		AddUserSum: userNumber.UserCount,
	}, nil
}
