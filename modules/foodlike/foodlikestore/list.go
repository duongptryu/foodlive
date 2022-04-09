package foodlikestore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodlike/foodlikemodel"
)

func (s sqlStore) ListUsersLikFood(ctx context.Context, conditions map[string]interface{}, filter *foodlikemodel.Filter, paging *common.Paging, moreKeys ...string) ([]foodlikemodel.FoodLike, error) {
	var result []foodlikemodel.FoodLike

	db := s.db

	db = db.Table(foodlikemodel.FoodLike{}.TableName()).Where(conditions)

	if v := filter; v != nil {

	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i, _ := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
