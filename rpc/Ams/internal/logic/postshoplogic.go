package logic

import (
	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"
	"DP/rpc/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostShopLogic {
	return &PostShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PostShop 店铺信息添加
func (l *PostShopLogic) PostShop(in *ams.PostShopReq) (*ams.PostShopResp, error) {

	//身份验证
	admin, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.AdminId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if admin == nil {
		return nil, errors.New("权限不足")
	}

	//准备插入数据
	data := model.Shop{
		ShopId:       in.ShopId,
		ShopName:     in.ShopName,
		ShopAddress:  in.ShopAddress,
		ShopCity:     in.ShopCity,
		ShopAdmin:    "",
		CreationTime: time.Now(),
	}

	//插入数据
	_, err = l.svcCtx.ShopModel.TransactInsertShop(l.ctx, nil, &data)
	if err != nil {
		fmt.Println("----+++++")
		return nil, err
	}

	return &ams.PostShopResp{
		Code:   200,
		Msg:    "设置成功",
		ShopId: in.ShopId,
	}, nil
}
