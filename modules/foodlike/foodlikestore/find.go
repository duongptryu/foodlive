package foodlikestore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodlike/foodlikemodel"
)

func (s sqlStore) FindUserLikeFood(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*foodlikemodel.FoodLike, error) {
	db := s.db

	db = db.Table(foodlikemodel.FoodLike{}.TableName()).Where(condition)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result foodlikemodel.FoodLike

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
