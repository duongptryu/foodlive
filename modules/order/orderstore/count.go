package orderstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"time"
)

func (s *sqlStore) CountOrder(ctx context.Context, condition map[string]interface{}, conditionTime *time.Time) (int, error) {
	db := s.db

	var result int64

	db = db.Table(ordermodel.Order{}.TableName()).Where(condition)

	if conditionTime != nil {
		db = db.Where("created_at > ?", conditionTime)
	}

	if err := db.Count(&result).Error; err != nil {
		return 0, common.ErrDB(err)
	}

	return int(result), nil
}
