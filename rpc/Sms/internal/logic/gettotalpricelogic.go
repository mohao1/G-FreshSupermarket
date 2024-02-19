package logic

import (
	"DP/rpc/Sms/internal/svc"
	"DP/rpc/Sms/pb/sms"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTotalPriceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTotalPriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTotalPriceLogic {
	return &GetTotalPriceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetTotalPrice 计算总价
func (l *GetTotalPriceLogic) GetTotalPrice(in *sms.GetTotalPriceReq) (*sms.GetTotalPriceResp, error) {

	if len(in.ProductQuantity) != len(in.ProductId) || len(in.ProductQuantity) == 0 || len(in.ProductId) == 0 {
		return &sms.GetTotalPriceResp{
			Code:       "400",
			TotalPrice: "",
			Msg:        "数据错误",
		}, nil
	}

	TotalNumber := 0.0
	for i, id := range in.ProductId {
		if id[:3] == "LPR" {
			findOne, err := l.svcCtx.LowProductList.FindOne(l.ctx, id)
			if err != nil {
				return nil, err
			}
			if in.ProductQuantity[i] > findOne.Quota || in.ProductQuantity[i] > findOne.ProductQuantity {
				return &sms.GetTotalPriceResp{
					Code:       "400",
					TotalPrice: "",
					Msg:        "数据错误-购买数量超出可购买的数量",
				}, nil
			}
			float, _ := strconv.ParseFloat(findOne.Price, 64)
			Num := float * float64(in.ProductQuantity[i])
			TotalNumber += Num
		} else {
			findOne, err := l.svcCtx.ProductList.FindOne(l.ctx, id)
			if err != nil {
				return nil, err
			}
			if in.ProductQuantity[i] > findOne.ProductQuantity {
				return &sms.GetTotalPriceResp{
					Code:       "400",
					TotalPrice: "",
					Msg:        "数据错误-购买数量超出可购买的数量",
				}, nil
			}
			float, _ := strconv.ParseFloat(findOne.Price, 64)
			Num := float * float64(in.ProductQuantity[i])
			TotalNumber += Num
		}
	}
	return &sms.GetTotalPriceResp{
		Code:       "200",
		TotalPrice: strconv.FormatFloat(TotalNumber, 'f', -1, 64),
		Msg:        "计算完成",
	}, nil
}
