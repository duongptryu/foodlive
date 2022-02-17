package ordertrackingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) CreateOrderTracking(ctx context.Context, data *ordertrackingmodel.OrderTrackingCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
