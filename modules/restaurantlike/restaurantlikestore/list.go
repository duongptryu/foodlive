package restaurantlikestore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
)

func (s sqlStore) GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id"`
		LikeCount    int `gorm:"column:count"`
	}

	var listLike []sqlData

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).Select("restaurant_id, count(restaurant_id) as count").Where("restaurant_id in (?)", ids).Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
}

func (s sqlStore) GetUsersLikeRestaurant(ctx context.Context, conditions map[string]interface{}, filter *restaurantlikemodel.Filter, paging *common.Paging, moreKeys ...string) ([]common.SimpleUser, error) {
	var result []restaurantlikemodel.Like

	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("User")

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User
	}

	return users, nil
}

func (s *sqlStore) ListMyLikeRestaurant(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]restaurantlikemodel.MyRstLike, error) {
	var result []restaurantlikemodel.MyRstLike

	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(condition)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := filter; v != nil {

	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Order("created_at desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
