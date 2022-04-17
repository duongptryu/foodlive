package citystore

import (
	"context"
	"foodlive/modules/city/citymodel"
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

type CityStore interface {
	ListCity(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) ([]citymodel.City, error)
}
