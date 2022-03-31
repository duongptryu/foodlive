package restaurantstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
	"time"
)

func (s *sqlStore) CountRst(ctx context.Context, condition map[string]interface{}, conditionTime *time.Time) (int, error) {
	db := s.db

	var result int64

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition)

	if conditionTime != nil {
		db = db.Where("created_at > ?", conditionTime)
	}

	if err := db.Count(&result).Error; err != nil {
		return 0, common.ErrDB(err)
	}

	return int(result), nil
}
