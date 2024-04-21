package AdminDataSum

import (
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNewUserSumToDayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNewUserSumToDayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewUserSumToDayLogic {
	return &GetNewUserSumToDayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNewUserSumToDayLogic) GetNewUserSumToDay(req *types.GetNewUserSumToDayReq) (resp *types.AmsDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
