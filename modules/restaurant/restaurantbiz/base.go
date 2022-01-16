package restaurantbiz

import (
	"context"
	"foodlive/modules/city/citymodel"
)

type CityStore interface {
	FindCity(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*citymodel.City, error)
}
