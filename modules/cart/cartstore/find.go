package cartstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/cart/cartmodel"
)

func (s *sqlStore) FindCartItem(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*cartmodel.CartItem, error) {
	db := s.db

	db = db.Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result cartmodel.CartItem

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
