package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAdminLogic {
	return &DeleteAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteAdmin 删除可用管理账号
func (l *DeleteAdminLogic) DeleteAdmin(in *ams.DeleteAdminReq) (*ams.DeleteAdminResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//判断账号是否存在
		staff, err2 := l.svcCtx.StaffModel.TransactSelectStaff(ctx, session, in.StaffId)
		if err2 != nil {
			return err2
		}
		if staff == nil {
			return errors.New("用户的账号不存在")
		}
		//判断是否已经管理商店
		if staff.ShopId != "" {
			return errors.New("用户的账号已经存在管理商店")
		}
		//判断权限
		if staff.PositionId != "3" {
			return errors.New("用户的账号权限不足")
		}

		//删除账号
		err2 = l.svcCtx.StaffModel.TransactDeleteStaff(l.ctx, session, staff.StaffId)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &ams.DeleteAdminResp{
		Code:    200,
		Msg:     "删除成功",
		StaffId: in.StaffId,
	}, nil
}
