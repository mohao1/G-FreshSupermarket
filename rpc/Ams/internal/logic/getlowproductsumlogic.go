package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLowProductSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLowProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLowProductSumLogic {
	return &GetLowProductSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLowProductSum 统计折扣商品总量
func (l *GetLowProductSumLogic) GetLowProductSum(in *ams.GetLowProductSumReq) (*ams.GetLowProductSumResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//获取数据
	productSum, err := l.svcCtx.LowProductListModel.SelectLowProductSum(l.ctx)
	if err != nil {
		return nil, err
	}

	//返回数据
	return &ams.GetLowProductSumResp{
		Code:       200,
		Msg:        "获取成功",
		ProductSum: productSum.CountNumber,
	}, nil
}
