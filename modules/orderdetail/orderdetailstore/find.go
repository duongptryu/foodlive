package orderdetailstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/orderdetail/orderdetailmodel"
)

func (s *sqlStore) FindOrderDetail(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) ([]orderdetailmodel.OrderDetail, error) {
	db := s.db

	db = db.Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result []orderdetailmodel.OrderDetail

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
