package StoreServices

import (
	"DP/rpc/Sms/smsclient"
	"context"

	"DP/DP/internal/svc"
	"DP/DP/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnnouncementListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAnnouncementListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnnouncementListLogic {
	return &GetAnnouncementListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAnnouncementListLogic) GetAnnouncementList(req *types.AnnouncementListReq) (resp *types.DataResp, err error) {

	//调用RPC的服务
	announcementList, err := l.svcCtx.SmsRpcClient.GetAnnouncementList(l.ctx, &smsclient.AnnouncementListReq{
		ShopId: req.ShopId,
	})

	//错误数据处理
	if err != nil {
		return &types.DataResp{
			Code: 400,
			Msg:  "请求错误err:" + err.Error(),
			Data: nil,
		}, nil
	}

	//正确数据返回
	return &types.DataResp{
		Code: 200,
		Msg:  "数据请求成功",
		Data: announcementList,
	}, nil
}
