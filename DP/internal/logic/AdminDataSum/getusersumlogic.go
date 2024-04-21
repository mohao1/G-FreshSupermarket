package AdminDataSum

import (
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSumLogic {
	return &GetUserSumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserSumLogic) GetUserSum(req *types.GetUserSumReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
