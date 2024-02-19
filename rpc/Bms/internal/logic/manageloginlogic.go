package logic

import (
	"DP/rpc/Utile"
	"context"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManageLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewManageLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManageLoginLogic {
	return &ManageLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ManageLogin 商户管理员的登录
func (l *ManageLoginLogic) ManageLogin(in *bms.ManageLoginReq) (*bms.ManageLoginResp, error) {

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
		return &bms.ManageLoginResp{
			Code:        400,
			Msg:         "账号权限不足",
			AccessToken: "",
		}, nil
	}

	if staff.Password == Utile.StrMD5ByStr(in.Password) {
		token, _ := Utile.GetJWTToken(staff.StaffId)
		return &bms.ManageLoginResp{
			Code:        200,
			Msg:         "登录成功",
			AccessToken: token,
		}, nil
	} else {
		return &bms.ManageLoginResp{
			Code:        400,
			Msg:         "密码错误登录失败",
			AccessToken: "",
		}, nil
	}
}
