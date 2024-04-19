package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &ams.GetUserSumResp{}, nil
}
