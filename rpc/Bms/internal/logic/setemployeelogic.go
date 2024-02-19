package logic

import (
	"DP/rpc/model"
	"context"
	"time"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetEmployeeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetEmployeeLogic {
	return &SetEmployeeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SetEmployee 设置员工
func (l *SetEmployeeLogic) SetEmployee(in *bms.SetEmployeeReq) (*bms.SetEmployeeResp, error) {

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
		return &bms.SetEmployeeResp{
			Code:    400,
			Msg:     "账号权限不足",
			StaffId: "",
		}, nil
	}

	findOne, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.NewStaffId)

	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if findOne != nil {
		return &bms.SetEmployeeResp{
			Code:    400,
			Msg:     "员工已经存在",
			StaffId: in.NewStaffId,
		}, nil
	}

	_, err = l.svcCtx.StaffModel.Insert(l.ctx, &model.Staff{
		StaffId:      in.NewStaffId,
		StaffName:    in.StaffName,
		PositionId:   in.PositionId,
		Password:     in.PassWord,
		ShopId:       staff.ShopId,
		CreationTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &bms.SetEmployeeResp{
		Code:    200,
		Msg:     "创建成功",
		StaffId: in.NewStaffId,
	}, nil
}
