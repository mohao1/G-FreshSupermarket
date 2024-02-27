package logic

import (
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminLogic {
	return &GetAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetAdmin 查看可用管理账号
func (l *GetAdminLogic) GetAdmin(in *ams.GetAdminReq) (*ams.GetAdminResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//查看可用管理账户列表
	staffs, err := l.svcCtx.StaffModel.TransactSelectStaffList(l.ctx, nil, "")
	if err != nil {
		return nil, err
	}

	//封装返回数据
	staffList := make([]*ams.ShopAdmin, len(*staffs))

	for i, staff := range *staffs {
		staffList[i] = &ams.ShopAdmin{
			StaffId:      staff.StaffId,
			StaffName:    staff.StaffName,
			PositionName: staff.PositionId,
			CreationTime: staff.CreationTime.String(),
		}
	}

	return &ams.GetAdminResp{
		Code:      200,
		Msg:       "获取成功",
		ShopAdmin: staffList,
	}, nil
}
