package logic

import (
	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"
	"DP/rpc/Utile"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *sms.LoginReq) (*sms.LoginResp, error) {
	if in.Phone == "" || in.Password == "" {
		return &sms.LoginResp{
			Code:        400,
			AccessToken: "",
			Msg:         "输入信息错误",
		}, nil
	}

	user, _ := l.svcCtx.UserModel.PhoneSelectUser(l.ctx, nil, in.Phone)
	if user == nil {
		return &sms.LoginResp{
			Code:        400,
			AccessToken: "",
			Msg:         "用户的数据不存在",
		}, nil
	} else if user.Password != Utile.StrMD5ByStr(in.Password) {
		return &sms.LoginResp{
			Code:        400,
			AccessToken: "",
			Msg:         "密码错误",
		}, nil
	} else {
		token, _ := Utile.GetJWTToken(user.Id)
		return &sms.LoginResp{
			Code:        200,
			AccessToken: token,
			Msg:         "登录成功",
		}, nil
	}
}
