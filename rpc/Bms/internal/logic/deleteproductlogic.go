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

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProduct 删除发布普通商品
func (l *DeleteProductLogic) DeleteProduct(in *bms.DeleteProductReq) (*bms.DeleteProductResp, error) {

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

	//上锁
	l.rWMutex.Lock()
	defer l.rWMutex.Unlock()

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
			ProductName:     product.ProductName,
			ProductTitle:    product.ProductTitle,
			ProductTypeId:   product.ProductTypeId,
			ProductQuantity: product.ProductQuantity,
			ProductPicture:  product.ProductPicture,
			Price:           product.Price,
			ProductSize:     product.ProductSize,
			ShopId:          product.ShopId,
			Producer:        product.Producer,
			DeleteKey:       1,
			CreationTime:    product.CreationTime,
			UpdataTime:      time.Now(),
		}

		//更新数据(虚拟删除操作)
		err2 = l.svcCtx.ProductListModel.TransactUpdateProductData(ctx, session, &data)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &bms.DeleteProductResp{
		Code:      200,
		Msg:       "删除成功",
		ProductId: in.ProductId,
	}, nil
}
