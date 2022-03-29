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
	FindOwnerRestaurant(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*restaurantownermodel.OwnerRestaurant, error)
	UpdateOwnerRestaurant(ctx context.Context, id int, data *restaurantownermodel.OwnerRestaurantUpdate) error
}