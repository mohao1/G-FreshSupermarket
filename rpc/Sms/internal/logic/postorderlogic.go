package logic

import (
	"DP/rpc/model"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"sync"
	"time"

	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RWMutex sync.RWMutex
}

func NewPostOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostOrderLogic {
	return &PostOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PostOrder 下单
func (l *PostOrderLogic) PostOrder(in *sms.PostOrderReq) (*sms.PostOrderResp, error) {
	if len(in.ProductId) == 0 || len(in.ProductQuantity) == 0 || len(in.ProductQuantity) != len(in.ProductId) {
		return &sms.PostOrderResp{
			Code:        400,
			OrderNumber: "",
			Msg:         "数据错误",
		}, nil
	}

	//计算订单
	TotalNumber := 0.0
	var Number int64 = 0
	for i, product := range in.ProductId {
		Number += in.ProductQuantity[i]
		if product[:3] == "LPR" {
			findOne, err := l.svcCtx.LowProductList.FindOne(l.ctx, product)
			if err != nil {
				return nil, err
			}
			if in.ProductQuantity[i] > findOne.Quota || in.ProductQuantity[i] > findOne.ProductQuantity {
				return &sms.PostOrderResp{
					Code:        400,
					OrderNumber: "",
					Msg:         "数据错误-购买数量超出可购买的数量",
				}, nil
			}
			float, _ := strconv.ParseFloat(findOne.Price, 64)
			Num := float * float64(in.ProductQuantity[i])
			TotalNumber += Num
		} else {
			findOne, err := l.svcCtx.ProductList.FindOne(l.ctx, product)
			if err != nil {
				return nil, err
			}
			if in.ProductQuantity[i] > findOne.ProductQuantity {
				return &sms.PostOrderResp{
					Code:        400,
					OrderNumber: "",
					Msg:         "数据错误-购买数量超出可购买的数量",
				}, nil
			}
			float, _ := strconv.ParseFloat(findOne.Price, 64)
			Num := float * float64(in.ProductQuantity[i])
			TotalNumber += Num
		}
	}
	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()
	//生成订单编号
	OrderNumber := uuid.New().String()
	err := l.svcCtx.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		for i, id := range in.ProductId {
			//锁定数量减少数量并且生成对应订单
			err := l.placeOrder(in.ProductQuantity[i], id, OrderNumber, session)
			if err != nil {
				return err
			}
		}

		deliveryTime, _ := strconv.Atoi(in.DeliveryTime)

		//生成总的订单
		data := model.OrderNumber{
			OrderNumber:       OrderNumber,
			CustomerId:        in.UserId,
			TotalPrice:        strconv.Itoa(int(TotalNumber)),
			Total:             Number,
			Payment:           0,
			ShopId:            in.ShopId,
			OrderOver:         0,
			OrderReceive:      0,
			ConfirmedDelivery: 0,
			CreationTime:      time.Now(),
			DeliveryTime:      time.Unix(int64(deliveryTime), 0),
			ConfirmTime:       sql.NullTime{},
			UpdataTime:        time.Now(),
			EndTime:           sql.NullTime{},
			DeleteKey:         0,
			Notes:             in.Notes,
		}

		_, err := l.svcCtx.OrderNumberModel.TransactInsert(ctx, session, &data)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &sms.PostOrderResp{
		Code:        200,
		OrderNumber: OrderNumber,
		Msg:         "下单成功",
	}, nil
}

func (l *PostOrderLogic) placeOrder(i int64, id string, OrderNumber string, session sqlx.Session) error {
	var order model.Order
	orderId := uuid.New().String()
	//扣除库存数量
	if id[:3] == "LPR" {
		lowProduct, err := l.svcCtx.LowProductList.TransactSelectLowProductList(l.ctx, session, id)
		if err != nil {
			return err
		}
		quantity := lowProduct.ProductQuantity
		if quantity < int64(i) {
			return errors.New("数量不足")
		}
		err = l.svcCtx.LowProductList.TransactUpDataLowProductList(l.ctx, session, id, int(quantity-i))
		if err != nil {
			return err
		}
		atoi, err := strconv.Atoi(lowProduct.Price)
		if err != nil {
			return err
		}

		order = model.Order{
			OrderId:        orderId,
			OrderName:      lowProduct.ProductName,
			OrderTitle:     lowProduct.ProductTitle,
			Price:          strconv.Itoa(atoi * int(i)),
			ProductTypeId:  lowProduct.ProductTypeId,
			OrderQuantity:  i,
			ProductSize:    lowProduct.ProductSize,
			ProductPicture: lowProduct.ProductPicture,
			OrderNumber:    OrderNumber,
			ShopId:         lowProduct.ShopId,
			CreationTime:   time.Now(),
			UpdataTime:     time.Now(),
			DeleteKey:      0,
		}
	} else {
		product, err := l.svcCtx.ProductList.TransactSelectProductList(l.ctx, session, id)
		if err != nil {
			return err
		}
		quantity := product.ProductQuantity
		if quantity < int64(i) {
			return errors.New("数量不足")
		}
		err = l.svcCtx.ProductList.TransactUpDataProductList(l.ctx, session, id, int(quantity-i))
		if err != nil {
			return err
		}

		atoi, err := strconv.Atoi(product.Price)
		if err != nil {
			return err
		}
		order = model.Order{
			OrderId:        orderId,
			OrderName:      product.ProductName,
			OrderTitle:     product.ProductTitle,
			Price:          strconv.Itoa(atoi * int(i)),
			ProductTypeId:  product.ProductTypeId,
			OrderQuantity:  i,
			ProductSize:    product.ProductSize,
			ProductPicture: product.ProductPicture,
			OrderNumber:    OrderNumber,
			ShopId:         product.ShopId,
			CreationTime:   time.Now(),
			UpdataTime:     time.Now(),
			DeleteKey:      0,
		}
	}

	_, err := l.svcCtx.OrderModel.TransactInsert(l.ctx, session, &order)

	if err != nil {
		return err
	}

	return nil
}
