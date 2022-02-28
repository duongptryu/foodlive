package orderdetailstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/orderdetail/orderdetailmodel"
)

func (s *sqlStore) CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetailCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) CreateBulkOrderDetail(ctx context.Context, data []orderdetailmodel.OrderDetailCreate) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
