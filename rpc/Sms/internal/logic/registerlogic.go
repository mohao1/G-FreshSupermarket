package logic

import (
	"DP/rpc/Utile"
	"DP/rpc/model"
	"context"
	"fmt"
	"github.com/google/uuid"

	"time"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *sms.RegisterReq) (*sms.RegisterResp, error) {
	user, err := l.svcCtx.UserModel.PhoneSelectUser(l.ctx, nil, in.Phone)
	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Println(err.Error())
		return nil, err
	}
	if user != nil {
		return &sms.RegisterResp{
			Code: 400,
			Data: "用户已经存在",
		}, nil
	}

	//生成UUID唯一主键
	NewUUID := uuid.New().String()
	//MD5加密Pwd
	pwd := Utile.StrMD5ByStr(in.Password)
	NewUser := &model.User{
		Id:               NewUUID,
		Name:             in.Name,
		Phone:            in.Phone,
		Password:         pwd,
		PositionId:       "1",
		RegistrationTime: time.Now(),
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, NewUser)
	if err != nil {
		return nil, err
	}

	return &sms.RegisterResp{
		Code: 200,
		Data: "注册成功",
	}, nil
}
