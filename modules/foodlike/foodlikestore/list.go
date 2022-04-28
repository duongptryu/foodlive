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

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Find(&result).Order("id").Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (s sqlStore) GetFoodLiked(ctx context.Context, ids []int, userId int) (map[int]bool, error) {
	result := make(map[int]bool)

	type sqlData struct {
		FoodId int `gorm:"column:food_id"`
	}

	var listLike []sqlData

	if err := s.db.Table(foodlikemodel.FoodLike{}.TableName()).Select("food_id").Where(map[string]interface{}{"food_id": ids, "user_id": userId}).Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.FoodId] = true
	}

	return result, nil
}
