package logic

import (
	"context"
	"fmt"
	"strconv"

	"DP/rpc/Ams/internal/svc"
	"DP/rpc/Ams/pb/ams"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopAllStaffSumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopAllStaffSumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopAllStaffSumListLogic {
	return &GetShopAllStaffSumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShopAllStaffSumList 各个门店人数统计
func (l *GetShopAllStaffSumListLogic) GetShopAllStaffSumList(in *ams.GetShopAllStaffSumListReq) (*ams.GetShopAllStaffSumListResp, error) {

	staffCounts, err := l.svcCtx.StaffModel.SelectStaffCount(l.ctx)
	if err != nil {
		return nil, err
	}

	//准备数据容器
	staffCountList := make([]*ams.StaffSum, len(*staffCounts))

	//填充数据数值
	for i, staffCount := range *staffCounts {

		fmt.Println(staffCount.CountNumber)
		fmt.Println(staffCount.ShopId)
		shop, err2 := l.svcCtx.ShopModel.FindOne(l.ctx, staffCount.ShopId)
		if err2 != nil {
			return nil, err2
		}

		staffCountList[i] = &ams.StaffSum{
			ShopId:      staffCount.ShopId,
			ShopName:    shop.ShopName,
			StaffNumber: strconv.Itoa(int(staffCount.CountNumber)),
		}
	}

	return &ams.GetShopAllStaffSumListResp{
		Code:         200,
		Msg:          "获取成功",
		StaffSumList: staffCountList,
	}, nil
}
