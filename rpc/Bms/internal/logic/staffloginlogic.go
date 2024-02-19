package logic

import (
	"DP/rpc/Utile"
	"context"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type StaffLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStaffLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StaffLoginLogic {
	return &StaffLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// StaffLogin 普通员工登录
func (l *StaffLoginLogic) StaffLogin(in *bms.StaffLoginReq) (*bms.StaffLoginResp, error) {

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

	if position.PositionName != "员工" {
		return &bms.StaffLoginResp{
			Code:        400,
			Msg:         "账号权限不足",
			AccessToken: "",
		}, nil
	}

	if staff.Password == Utile.StrMD5ByStr(in.Password) {
		token, _ := Utile.GetJWTToken(staff.StaffId)
		return &bms.StaffLoginResp{
			Code:        200,
			Msg:         "登录成功",
			AccessToken: token,
		}, nil
	} else {
		return &bms.StaffLoginResp{
			Code:        400,
			Msg:         "密码错误登录失败",
			AccessToken: "",
		}, nil
	}
}
