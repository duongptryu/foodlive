package orderdetailstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/orderdetail/orderdetailmodel"
)

func (s *sqlStore) UpdateOrderDetail(ctx context.Context, id int, data *orderdetailmodel.OrderDetailUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
