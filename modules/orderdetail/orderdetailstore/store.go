package orderdetailstore

import (
	"context"
	"foodlive/modules/orderdetail/orderdetailmodel"
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
	CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetailCreate) error
	UpdateOrderDetail(ctx context.Context, id int, data *orderdetailmodel.OrderDetailUpdate) error
	CreateBulkOrderDetail(ctx context.Context, data []orderdetailmodel.OrderDetailCreate) error
}
