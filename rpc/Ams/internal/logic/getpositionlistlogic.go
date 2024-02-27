package logic

import (
	"context"
	"errors"
	"strconv"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

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

// GetPositionList 身份信息列表查看
func (l *GetPositionListLogic) GetPositionList(in *ams.GetPositionListReq) (*ams.GetPositionListResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查询身份信息列表
	list, err := l.svcCtx.PositionModel.SelectPositionList(l.ctx)
	if err != nil {
		return nil, err
	}

	//准备数据信息
	positionList := make([]*ams.PositionData, len(*list))

	for i, position := range *list {
		positionList[i] = &ams.PositionData{
			PositionId:    position.Id,
			PositionName:  position.PositionName,
			PositionGrade: strconv.Itoa(int(position.Grade)),
		}
	}

	return &ams.GetPositionListResp{
		Code:             200,
		Msg:              "获取成功",
		PositionDataList: positionList,
	}, nil
}
