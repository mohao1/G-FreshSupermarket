package logic

import (
	"DP/rpc/Utile"
	"DP/rpc/model"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"sync"
	"time"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReleaseLowProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewReleaseLowProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseLowProductLogic {
	return &ReleaseLowProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReleaseLowProduct 设置发布折扣商品
func (l *ReleaseLowProductLogic) ReleaseLowProduct(in *bms.ReleaseLowProductReq) (*bms.ReleaseLowProductResp, error) {

	//查询店铺信息和个人的信息
	staff, err := l.svcCtx.StaffModel.FindOne(l.ctx, in.StaffId)
	if err != nil {
		return nil, err
	}

	//进行权限判断
	position, err := l.svcCtx.PositionModel.FindOne(l.ctx, staff.PositionId)
	if err != nil {
		return nil, err
	}
	if position.PositionName != "经理" {
		return nil, errors.New("权限不足")
	}

	//判断商品类型是否存在类型列表之中
	productType, err := l.svcCtx.ProductTypeModel.FindOne(l.ctx, in.ProductTypeId)
	if err != nil {
		return nil, err
	}
	if productType == nil {
		return &bms.ReleaseLowProductResp{
			Code:      400,
			Msg:       "商品类型数据错误",
			ProductId: "",
		}, nil
	}

	//上锁
	l.rWMutex.Lock()
	defer l.rWMutex.Unlock()

	//进行插入
	//生成商品Id
	ProductId := "LPR" + Utile.GetUUIDBy8()

	data := model.LowProductList{
		ProductId:       ProductId,
		ProductName:     in.ProductName,
		ProductTitle:    in.ProductTitle,
		ProductTypeId:   in.ProductTypeId,
		ProductQuantity: in.ProductQuantity,
		ProductPicture:  in.ProductPicture,
		Price:           in.Price,
		Producer:        in.Producer,
		Quota:           in.Quota,
		ProductSize:     in.ProductSize,
		ShopId:          staff.ShopId,
		DeleteKey:       0,
		CreationTime:    time.Now(),
		UpdataTime:      time.Now(),
		StartTime:       time.Unix(in.StartTime, 0),
		EndTime:         time.Unix(in.StartTime, 0),
	}

	sdata := model.SalesRecords{
		SalesRecordsId: uuid.New().String(),
		ProductId:      ProductId,
		SalesQuantity:  0,
		TotalPrice:     "0",
		ShopId:         staff.ShopId,
		CreationTime:   time.Now(),
		UpdataTime:     time.Now(),
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		//插入商品数据
		_, err = l.svcCtx.LowProductListModel.TransactInsert(ctx, session, &data)
		if err != nil {
			return err
		}

		//插入报表数据
		_, err2 := l.svcCtx.SalesRecordsModel.TransactInsert(ctx, session, &sdata)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &bms.ReleaseLowProductResp{
		Code:      200,
		Msg:       "设置成功",
		ProductId: ProductId,
	}, nil
}
