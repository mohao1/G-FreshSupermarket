package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderSumLogic {
	return &GetOrderSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderSum 今日消费用户数量
func (l *GetOrderSumLogic) GetOrderSum(in *ams.GetOrderSumReq) (*ams.GetOrderSumResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询信息
	OrderSum, err := l.svcCtx.OrderNumberModel.SelectOrderTheDaySum(l.ctx)
	if err != nil {
		return nil, err
	}

	return &ams.GetOrderSumResp{
		OrderSum: OrderSum.OrderCount,
	}, nil
}
