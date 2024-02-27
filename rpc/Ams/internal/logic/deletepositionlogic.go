package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePositionLogic {
	return &DeletePositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeletePosition 身份信息删除
func (l *DeletePositionLogic) DeletePosition(in *ams.DeletePositionReq) (*ams.DeletePositionResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		//查询是否存在
		position, err2 := l.svcCtx.PositionModel.TransactSelectPosition(ctx, session, in.PositionId)
		if err2 != nil {
			return err2
		}
		if position == nil {
			return errors.New("修改的类型不存在")
		}

		//删除数据
		err2 = l.svcCtx.PositionModel.TransactDeletePosition(ctx, session, in.PositionId)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &ams.DeletePositionResp{
		Code:       200,
		Msg:        "删除成功",
		PositionId: in.PositionId,
	}, nil
}
