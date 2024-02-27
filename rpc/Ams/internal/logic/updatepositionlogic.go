package logic

import (
	"DP/rpc/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDatePositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpDatePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDatePositionLogic {
	return &UpDatePositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpDatePosition 身份信息修改
func (l *UpDatePositionLogic) UpDatePosition(in *ams.UpDatePositionReq) (*ams.UpDatePositionResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//开启事务
	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		//查询是否存在
		position, err2 := l.svcCtx.PositionModel.TransactSelectPosition(ctx, session, in.PositionId)
		if err2 != nil {
			return err2
		}
		if position == nil {
			return errors.New("修改的类型不存在")
		}

		//准备数据
		data := model.Position{
			Id:           in.PositionId,
			PositionName: in.PositionName,
			Grade:        in.PositionGrade,
		}

		//修改数据
		err2 = l.svcCtx.PositionModel.TransactUpDatePosition(ctx, session, &data)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &ams.UpDatePositionResp{
		Code:       200,
		Msg:        "修改成功",
		PositionId: in.PositionId,
	}, nil
}
