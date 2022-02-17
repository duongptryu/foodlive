package orderstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
)

func (s *sqlStore) FindOrder(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*ordermodel.Order, error) {
	db := s.db

	db = db.Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result ordermodel.Order

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}