package restaurantownerstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantowner/restaurantownermodel"
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

type RestaurantOwnerStore interface {
	ListOwnerRestaurant(ctx context.Context, condition map[string]interface{}, filter *restaurantownermodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantownermodel.OwnerRestaurant, error)
}