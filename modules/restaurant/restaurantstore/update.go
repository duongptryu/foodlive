package restaurantstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UpdateRestaurantStatus(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdateStatus) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UpdateRestaurantRating(ctx context.Context, rstId int, rating float64) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", rstId).Update("rating", rating).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseRatingCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Update("rating_count", gorm.Expr("rating_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
