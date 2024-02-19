package User

import (
	"DP/rpc/Sms/smsclient"
	"context"
	"regexp"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterReq) (resp *types.DataResp, err error) {
	if !l.checkPhone(req.Phone) {
		return &types.DataResp{
			Code: 400,
			Msg:  "手机号不符合要求",
			Data: nil,
		}, nil
	} else {
		register, err2 := l.svcCtx.SmsRpcClient.Register(l.ctx, &smsclient.RegisterReq{
			Name:     req.Name,
			Phone:    req.Phone,
			Password: req.Password,
		})
		if err2 != nil {
			return nil, err2
		}

		return &types.DataResp{
			Code: int(register.Code),
			Msg:  register.Data,
			Data: nil,
		}, nil
	}
}

func (l *RegisterLogic) checkPhone(phone string) bool {

	reg := `^1[3456789]\d{9}$`
	matchString, err := regexp.MatchString(reg, phone)
	if err != nil {
		return false
	}
	return matchString
}
