package AdminUpDate

import (
	"DP/rpc/Ams/amsclient"
	"context"
	"fmt"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDatePassWordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDatePassWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDatePassWordLogic {
	return &UpDatePassWordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpDatePassWord 系统管理人员修改密码
func (l *UpDatePassWordLogic) UpDatePassWord(req *types.UpDatePassWordReq) (resp *types.AmsDataResp, err error) {

	//获取JWT中的UserId
	AdminId := fmt.Sprint(l.ctx.Value("jwtUserId"))

	//准备数据
	updateLoginPassWordReq := amsclient.UpdateLoginPassWordReq{
		AdminName:   AdminId,
		PassWord:    req.PassWord,
		NewPassWord: req.NewPassWord,
	}

	//调用RPC的服务
	updateLoginPassWordResp, err := l.svcCtx.AmsRpcClient.AdminUpdateLoginPassWord(l.ctx, &updateLoginPassWordReq)
	if err != nil {
		return &types.AmsDataResp{
			Code: 400,
			Msg:  "调用错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	return &types.AmsDataResp{
		Code: updateLoginPassWordResp.Code,
		Msg:  updateLoginPassWordResp.Msg,
		Data: nil,
	}, nil
}
