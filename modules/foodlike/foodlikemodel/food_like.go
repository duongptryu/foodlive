package foodlikemodel

import "time"

type FoodLike struct {
	FoodId    int       `json:"food_id" gorm:"food_id"`
	UserId    int       `json:"user_id" gorm:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}

func (FoodLike) TableName() string {
	return "food_like"
}

type FoodLikeCreate struct {
	FoodId    int        `json:"food_id" gorm:"food_id"`
	UserId    int        `json:"user_id" gorm:"user_id"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
}

func (FoodLikeCreate) TableName() string {
	return FoodLike{}.TableName()
}
