package foodratingmodel

import "foodlive/common"

const (
	EntityName = "FoodRating"
)

type FoodRating struct {
	common.SQLModel
	UserId  int                `json:"user_id" gorm:"user_id"`
	User    *common.SimpleUser `json:"user" gorm:"preload:false"`
	FoodId  int                `json:"food_id" gorm:"food_id"`
	Point   float64            `json:"point" gorm:"point"`
	Comment string             `json:"comment" gorm:"comment"`
	Status  bool               `json:"status" gorm:"status"`
}

func (FoodRating) TableName() string {
	return "food_rating"
}

type FoodRatingCreate struct {
	common.SQLModelCreate
	UserId  int    `json:"-" gorm:"user_id"`
	FoodId  int    `json:"food_id" gorm:"food_id"`
	Point   int    `json:"point" gorm:"point"`
	Comment string `json:"comment" gorm:"comment"`
	Status  bool   `json:"-" gorm:"status"`
}

func (FoodRatingCreate) TableName() string {
	return FoodRating{}.TableName()
}

func (data *FoodRatingCreate) Validate() error {
	if data.Point < 0 || data.Point > 5 {
		return ErrPointMustIn0To5
	}
	data.Status = true
	return nil
}

func (data *FoodRatingCreate) GetFoodId() int {
	return data.FoodId
}

func (data *FoodRatingCreate) GetPoint() int {
	return data.Point
}

type FoodRatingUpdate struct {
	common.SQLModelUpdate
	Point   float64 `json:"point" gorm:"point"`
	Comment string  `json:"comment" gorm:"comment"`
	Status  bool    `json:"-" gorm:"status"`
}

func (FoodRatingUpdate) TableName() string {
	return FoodRating{}.TableName()
}

func (data *FoodRatingUpdate) Validate() error {
	if data.Point < 0 || data.Point > 5 {
		return ErrPointMustIn0To5
	}
	return nil
}

var ErrPointMustIn0To5 = common.NewCustomError(nil, "Point rating must be stay in 0 to 5, and is integer", "ErrPointMustIn0To5")
var ErrAlreadyRating = common.NewCustomError(nil, "User already rating this food", "ErrAlreadyRating")
