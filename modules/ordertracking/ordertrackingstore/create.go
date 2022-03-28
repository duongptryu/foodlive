package ordertrackingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/ordertracking/ordertrackingmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) CreateOrderTracking(ctx context.Context, data *ordertrackingmodel.OrderTrackingCreate) (*gorm.DB, error) {
	db := s.db.Begin()

	if err := db.Create(data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return db, nil
}
