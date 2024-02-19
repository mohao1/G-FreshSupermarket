package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdvertisementModel = (*customAdvertisementModel)(nil)

type (
	// AdvertisementModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdvertisementModel.
	AdvertisementModel interface {
		advertisementModel
	}

	customAdvertisementModel struct {
		*defaultAdvertisementModel
	}
)

// NewAdvertisementModel returns a model for the database table.
func NewAdvertisementModel(conn sqlx.SqlConn) AdvertisementModel {
	return &customAdvertisementModel{
		defaultAdvertisementModel: newAdvertisementModel(conn),
	}
}
