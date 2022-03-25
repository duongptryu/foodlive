package ordertrackingstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) FindOrderTracking(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*ordertrackingmodel.OrderTracking, error) {
	db := s.db

	db = db.Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result ordertrackingmodel.OrderTracking

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
