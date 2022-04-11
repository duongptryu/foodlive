package foodstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/food/foodmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateFood(ctx context.Context, id int, data *foodmodel.FoodUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) UpdateRating(ctx context.Context, foodId int, rating float64) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).Where("id = ?", foodId).Update("rating", rating).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseRatingCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).Where("id = ?", id).Update("rating_count", gorm.Expr("rating_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
