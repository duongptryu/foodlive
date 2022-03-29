package restaurantstore

import (
	"context"
	"foodlive/modules/restaurant/restaurantmodel"
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

type RestaurantStore interface {
	FindRestaurant(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}
