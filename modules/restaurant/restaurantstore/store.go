package restaurantstore

import (
	"context"
	"foodlive/modules/restaurant/restaurantmodel"
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

type RestaurantStore interface {
	FindRestaurant(ctx context.Context, condition map[string]interface{}, filter *restaurantmodel.Filter, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	CountRst(ctx context.Context, condition map[string]interface{}, conditionTime *time.Time) (int, error)
	IncreaseRatingCount(ctx context.Context, id int) error
}
