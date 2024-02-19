package logic

import (
	"DP/rpc/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"sync"
	"time"

	"DP/rpc/Bms/internal/svc"
	"DP/rpc/Bms/pb/bms"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateLowProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewUpDateLowProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateLowProductLogic {
	return &UpDateLowProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpDateLowProduct 修改发布折扣商品
func (l *UpDateLowProductLogic) UpDateLowProduct(in *bms.UpDateLowProductReq) (*bms.UpDateLowProductResp, error) {

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
		return &bms.UpDateLowProductResp{
			Code:      400,
			Msg:       "商品类型数据错误",
			ProductId: "",
		}, nil
	}

	//上锁
	l.rWMutex.Lock()
	defer l.rWMutex.Unlock()

	//开启事务
	err = l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		//查询商品是否存在
		lowProduct, err2 := l.svcCtx.LowProductListModel.TransactSelectLowProductList(ctx, session, in.ProductId)
		if err2 != nil {
			return err2
		}

		if lowProduct.DeleteKey != 0 {
			return errors.New("商品的数据不存在")
		}

		if lowProduct.ShopId != staff.ShopId {
			return errors.New("商店信息错误")
		}

		//定义数据
		data := model.LowProductList{
			ProductId:       in.ProductId,
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
			CreationTime:    lowProduct.CreationTime,
			UpdataTime:      time.Now(),
			StartTime:       time.Unix(in.StartTime, 0),
			EndTime:         time.Unix(in.EndTime, 0),
		}

		//修改数据
		err2 = l.svcCtx.LowProductListModel.TransactUpdateLowProductData(ctx, session, &data)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &bms.UpDateLowProductResp{
		Code:      200,
		Msg:       "修改成功",
		ProductId: in.ProductId,
	}, nil
}
