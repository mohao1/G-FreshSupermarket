package logic

import (
	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"
	"DP/rpc/Utile"
	"DP/rpc/model"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateLoginPassWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminUpdateLoginPassWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdateLoginPassWordLogic {
	return &AdminUpdateLoginPassWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AdminUpdateLoginPassWord 修改密码
func (l *AdminUpdateLoginPassWordLogic) AdminUpdateLoginPassWord(in *ams.UpdateLoginPassWordReq) (*ams.UpdateLoginPassWordResp, error) {

	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminName)
	if err != nil {
		return nil, err
	}

	if admin.Password != Utile.StrMD5ByStr(in.PassWord) {
		return nil, errors.New("输入的原密码错误")
	}

	if in.NewPassWord == in.PassWord {
		return nil, errors.New("新旧密码相同")
	}

	err = l.svcCtx.AdminModel.Update(l.ctx, &model.Admin{
		AdminId:  admin.AdminId,
		Password: Utile.StrMD5ByStr(in.NewPassWord),
	})
	if err != nil {
		return nil, err
	}

	return &ams.UpdateLoginPassWordResp{
		Code: 200,
		Msg:  "密码修改成功",
	}, nil
}
