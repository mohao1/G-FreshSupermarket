package User

import (
	"DP/rpc/Sms/smsclient"
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.DataResp, err error) {
	loginReq := smsclient.LoginReq{
		Phone:    req.UserName,
		Password: req.PassWord,
	}
	login, err := l.svcCtx.SmsRpcClient.Login(l.ctx, &loginReq)

	if err != nil {
		return nil, err
	}

	return &types.DataResp{
		Code: int(login.Code),
		Msg:  login.Msg,
		Data: login.AccessToken,
	}, nil
	return
}
