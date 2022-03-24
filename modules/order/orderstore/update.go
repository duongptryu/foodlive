package orderstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
)

func (s *sqlStore) UpdateOrder(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {
	db := s.db

	if err := db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
