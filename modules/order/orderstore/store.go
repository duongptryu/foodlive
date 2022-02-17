package orderstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
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
	CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) error
	FindOrder(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*ordermodel.Order, error)
	ListOrder(ctx context.Context,
		condition map[string]interface{},
		filter *ordermodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]ordermodel.Order, error)
	UpdateOrder(ctx context.Context, id int, data *ordermodel.OrderUpdate) error
}
