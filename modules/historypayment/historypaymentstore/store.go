package historypaymentstore

import (
	"context"
	"foodlive/modules/ordertracking/ordertrackingmodel"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type OrderStore interface {
	CreateOrderTracking(ctx context.Context, data *ordertrackingmodel.OrderTrackingCreate) error
	UpdateOrder(ctx context.Context, id int, data *ordertrackingmodel.OrderTrackingUpdate) error
}
