package orderstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"gorm.io/gorm"
)

func (s *sqlStore) CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) (*gorm.DB, error) {
	db := s.db.Begin()

	if err := db.Create(data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return db, nil
}
