package foodlikemodel

import (
	"foodlive/common"
	"foodlive/modules/food/foodmodel"
	"time"
)

type FoodLike struct {
	FoodId    int                `json:"food_id" gorm:"food_id"`
	Food      *foodmodel.Food    `json:"food" gorm:"preload:false"`
	UserId    int                `json:"user_id" gorm:"user_id"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false"`
	CreatedAt time.Time          `json:"created_at" gorm:"created_at"`
}

func (FoodLike) TableName() string {
	return "food_likes"
}

func (data *FoodLike) GetFoodId() int {
	return data.FoodId
}

type FoodLikeCreate struct {
	FoodId    int        `json:"food_id" gorm:"food_id"`
	UserId    int        `json:"-" gorm:"user_id"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
}

func (FoodLikeCreate) TableName() string {
	return FoodLike{}.TableName()
}

func (data *FoodLikeCreate) GetFoodId() int {
	return data.FoodId
}

var ErrUserAlreadyLikeFood = common.NewCustomError(nil, "User already like food", "ErrUserAlreadyLikeFood")
var ErrUserNotLikeFoodYet = common.NewCustomError(nil, "User not like food yet", "ErrUserNotLikeFoodYet")
