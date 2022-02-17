package orderstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
)

func (s *sqlStore) DeleteOrder (ctx context.Context, condition map[string]interface{}) error {
	db := s.db

	if err := db.Table(ordermodel.Order{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
} 