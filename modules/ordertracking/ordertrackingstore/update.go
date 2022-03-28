package ordertrackingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) UpdateOrderTracking(ctx context.Context, orderId int, data *ordertrackingmodel.OrderTrackingUpdate) error {
	db := s.db

	if err := db.Where("order_id = ?", orderId).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
