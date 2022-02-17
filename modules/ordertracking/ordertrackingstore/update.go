package ordertrackingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) UpdateOrder(ctx context.Context, id int, data *ordertrackingmodel.OrderTrackingUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
