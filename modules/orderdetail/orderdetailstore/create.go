package orderdetailstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/orderdetail/orderdetailmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetailCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) CreateBulkOrderDetail(ctx context.Context, data []orderdetailmodel.OrderDetailCreate) (*gorm.DB, error) {
	db := s.db.Begin()

	if err := db.Create(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return db, nil
}
