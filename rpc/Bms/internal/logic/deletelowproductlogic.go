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

type DeleteLowProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rWMutex sync.RWMutex
}

func NewDeleteLowProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLowProductLogic {
	return &DeleteLowProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteLowProduct 删除发布折扣商品
func (l *DeleteLowProductLogic) DeleteLowProduct(in *bms.DeleteLowProductReq) (*bms.DeleteLowProductResp, error) {

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
			ProductId:       lowProduct.ProductId,
			ProductName:     lowProduct.ProductName,
			ProductTitle:    lowProduct.ProductTitle,
			ProductTypeId:   lowProduct.ProductTypeId,
			ProductQuantity: lowProduct.ProductQuantity,
			ProductPicture:  lowProduct.ProductPicture,
			Price:           lowProduct.Price,
			Producer:        lowProduct.Producer,
			Quota:           lowProduct.Quota,
			ProductSize:     lowProduct.ProductSize,
			ShopId:          lowProduct.ShopId,
			DeleteKey:       1,
			CreationTime:    lowProduct.CreationTime,
			UpdataTime:      time.Now(),
			StartTime:       lowProduct.StartTime,
			EndTime:         lowProduct.EndTime,
		}

		//更新数据(虚拟删除操作)
		err2 = l.svcCtx.LowProductListModel.TransactUpdateLowProductData(ctx, session, &data)
		if err2 != nil {
			return err2
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &bms.DeleteLowProductResp{
		Code:      200,
		Msg:       "删除成功",
		ProductId: in.ProductId,
	}, nil
}
