package orderstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"gorm.io/gorm"
	"time"
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
	CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) (*gorm.DB, error)
	FindOrder(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*ordermodel.Order, error)
	ListOrder(ctx context.Context,
		condition map[string]interface{},
		filter *ordermodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]ordermodel.Order, error)
	UpdateOrder(ctx context.Context, id int, data *ordermodel.OrderUpdate) error
	CountOrder(ctx context.Context, condition map[string]interface{}, conditionTime *time.Time) (int, error)
	ListOrderWithoutPaging(ctx context.Context,
		condition map[string]interface{},
		filter *ordermodel.Filter,
		moreKey ...string,
	) ([]ordermodel.Order, error)
}
