package logic

import (
	"DP/rpc/Utile"
	"context"
	"errors"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AdminLogin 系统管理人员登录
func (l *AdminLoginLogic) AdminLogin(in *ams.AdminLoginReq) (*ams.AdminLoginResp, error) {

	//查询用户
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminName)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return &ams.AdminLoginResp{
			Code:        400,
			Msg:         "账号错误",
			AccessToken: "",
		}, nil
	}

	//校验密码
	if admin.Password != Utile.StrMD5ByStr(in.PassWord) {
		return nil, errors.New("密码错误")
	}

	//生成身份校验令牌
	token, err := Utile.GetJWTToken(admin.AdminId)
	if err != nil {
		return nil, err
	}

	return &ams.AdminLoginResp{
		Code:        200,
		Msg:         "登录成功",
		AccessToken: token,
	}, nil
}
