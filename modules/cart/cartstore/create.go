package cartstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/cart/cartmodel"
)

func (s *sqlStore) CreateCartItem(ctx context.Context, data *cartmodel.CartItemCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
