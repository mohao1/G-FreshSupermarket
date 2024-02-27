package logic

import (
	"DP/rpc/model"
	"context"
	"errors"
	"time"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostAdminLogic {
	return &PostAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PostAdmin 创建管理账号
func (l *PostAdminLogic) PostAdmin(in *ams.PostAdminReq) (*ams.PostAdminResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	data := model.Staff{
		StaffId:      in.StaffId,
		StaffName:    in.StaffName,
		PositionId:   "3",
		Password:     in.Password,
		ShopId:       "",
		CreationTime: time.Now(),
	}

	//创建账号
	_, err = l.svcCtx.StaffModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}

	return &ams.PostAdminResp{
		Code:    200,
		Msg:     "创建成功",
		StaffId: in.StaffId,
	}, nil
}
