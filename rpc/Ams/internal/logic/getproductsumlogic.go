package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSumLogic {
	return &GetProductSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductSum 统计普通商品总量
func (l *GetProductSumLogic) GetProductSum(in *ams.GetProductSumReq) (*ams.GetProductSumResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询信息
	productAllSum, err := l.svcCtx.ProductListModel.SelectProductAllSum(l.ctx)
	if err != nil {
		return nil, err
	}

	//返回数据
	return &ams.GetProductSumResp{
		Code:       200,
		Msg:        "获取成功",
		ProductSum: productAllSum.CountNumber,
	}, nil
}
