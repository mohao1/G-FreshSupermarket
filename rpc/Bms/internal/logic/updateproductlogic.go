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

type UpDateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewUpDateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateProductLogic {
	return &UpDateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpDateProduct 修改发布普通商品
func (l *UpDateProductLogic) UpDateProduct(in *bms.UpDateProductReq) (*bms.UpDateProductResp, error) {

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
		return &bms.UpDateProductResp{
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
		product, err2 := l.svcCtx.ProductListModel.TransactSelectProductList(ctx, session, in.ProductId)
		if err2 != nil {
			return err2
		}

		if product.DeleteKey != 0 {
			return errors.New("商品的数据不存在")
		}

		if product.ShopId != staff.ShopId {
			return errors.New("商店信息错误")
		}

		//定义数据
		data := model.ProductList{
			ProductId:       product.ProductId,
			ProductName:     in.ProductName,
			ProductTitle:    in.ProductTitle,
			ProductTypeId:   in.ProductTypeId,
			ProductQuantity: in.ProductQuantity,
			ProductPicture:  in.ProductPicture,
			Price:           in.Price,
			ProductSize:     in.ProductSize,
			ShopId:          staff.ShopId,
			Producer:        in.Producer,
			DeleteKey:       0,
			CreationTime:    product.CreationTime,
			UpdataTime:      time.Now(),
		}

		//修改数据
		err2 = l.svcCtx.ProductListModel.TransactUpdateProductData(ctx, session, &data)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &bms.UpDateProductResp{
		Code:      200,
		Msg:       "修改成功",
		ProductId: in.ProductId,
	}, nil
}
