package logic

import (
	"DP/rpc/Utile"
	"DP/rpc/model"
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostPositionLogic {
	return &PostPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PostPosition 身份信息设置
func (l *PostPositionLogic) PostPosition(in *ams.PostPositionReq) (*ams.PostPositionResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//生成随机Id
	positionId := Utile.GetUUIDBy8()

	//准备数据信息
	data := model.Position{
		Id:           positionId,
		PositionName: in.PositionName,
		Grade:        in.PositionGrade,
	}

	//插入数据
	_, err = l.svcCtx.PositionModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}

	return &ams.PostPositionResp{
		Code:       200,
		Msg:        "设置成功",
		PositionId: positionId,
	}, nil
}
