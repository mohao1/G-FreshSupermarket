package logic

import (
	"context"
	"strconv"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPositionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPositionListLogic {
	return &GetPositionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPositionList 获取身份列表
func (l *GetPositionListLogic) GetPositionList(in *bms.GetPositionListReq) (*bms.GetPositionListResp, error) {

	//查询账号是否存在
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//进行权限判断
	position, err := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
	if err != nil {
		return nil, err
	}

	if position.PositionName != "经理" {
		return &bms.GetPositionListResp{
			Code:     400,
			Msg:      "账号权限不足",
			Position: nil,
		}, nil
	}

	//进行数据列表查询
	positionList, err := l.svcCtx.PositionModel.SelectPositionList(l.ctx)
	if err != nil {
		return nil, err
	}

	positions := make([]*bms.Position, len(*positionList))

	for i, p := range *positionList {
		positions[i] = &bms.Position{
			PositionId:   p.Id,
			PositionName: p.PositionName,
			Grade:        strconv.Itoa(int(p.Grade)),
		}
	}

	return &bms.GetPositionListResp{
		Code:     200,
		Msg:      "获取成功",
		Position: positions,
	}, nil
}
