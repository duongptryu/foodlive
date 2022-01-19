package cartstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/cart/cartmodel"
)

func (s *sqlStore) DeleteCartItem(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db

	if err := db.Table(cartmodel.CartItem{}.TableName()).Where(conditions).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
