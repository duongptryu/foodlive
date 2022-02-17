package cartstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/cart/cartmodel"
)

func (s *sqlStore) UpdateCartItem(ctx context.Context, data *cartmodel.CartItemUpdate) error {
	db := s.db

	if err := db.Where("user_id = ? AND food_id = ?", data.UserId, data.FoodId).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
