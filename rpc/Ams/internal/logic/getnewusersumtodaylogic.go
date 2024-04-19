package logic

import (
	"context"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

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

// 今日新增用户数量
func (l *GetNewUserSumToDayLogic) GetNewUserSumToDay(in *ams.GetNewUserSumToDayReq) (*ams.GetNewUserSumToDayResp, error) {
	// todo: add your logic here and delete this line

	return &ams.GetNewUserSumToDayResp{}, nil
}
