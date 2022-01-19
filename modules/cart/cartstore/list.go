package cartstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/cart/cartmodel"
)

func (s *sqlStore) ListCartItem(ctx context.Context, condition map[string]interface{}, filter *cartmodel.Filter, moreKey ...string,
) ([]cartmodel.CartItem, error) {
	var result []cartmodel.CartItem

	db := s.db

	db = db.Table(cartmodel.CartItem{}.TableName()).Where(condition)

	if v := filter; v != nil {
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
