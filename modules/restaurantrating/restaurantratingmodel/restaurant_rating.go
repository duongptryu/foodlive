package restaurantratingmodel

import (
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

const (
	EntityName = "RestaurantRating"
)

type RestaurantRating struct {
	common.SQLModel
	UserId       int                         `json:"user_id" gorm:"user_id"`
	User         *common.SimpleUser          `json:"user" gorm:"preload:false"`
	RestaurantId int                         `json:"restaurant_id" gorm:"restaurant_id"`
	Restaurant   *restaurantmodel.Restaurant `json:"restaurant" gorm:"preload:false"`
	Point        float64                     `json:"point" gorm:"point"`
	Comment      string                      `json:"comment" gorm:"comment"`
	Status       bool                        `json:"status" gorm:"status"`
}

func (RestaurantRating) TableName() string {
	return "restaurant_ratings"
}

type RestaurantRatingCreate struct {
	common.SQLModelCreate
	UserId       int    `json:"-" gorm:"user_id"`
	RestaurantId int    `json:"restaurant_id" gorm:"restaurant_id"`
	Point        int    `json:"point" gorm:"point"`
	Comment      string `json:"comment" gorm:"comment"`
	Status       bool   `json:"-" gorm:"status"`
}

func (RestaurantRatingCreate) TableName() string {
	return RestaurantRating{}.TableName()
}

func (data *RestaurantRatingCreate) Validate() error {
	if data.Point < 0 || data.Point > 5 {
		return ErrPointMustIn0To5
	}
	data.Status = true
	return nil
}

func (data *RestaurantRatingCreate) GetRestaurantId() int {
	return data.RestaurantId
}

func (data *RestaurantRatingCreate) GetPoint() int {
	return data.Point
}

type RestaurantRatingUpdate struct {
	common.SQLModelUpdate
	Point   float64 `json:"point" gorm:"point"`
	Comment string  `json:"comment" gorm:"comment"`
	Status  bool    `json:"-" gorm:"status"`
}

func (RestaurantRatingUpdate) TableName() string {
	return RestaurantRating{}.TableName()
}

func (data *RestaurantRatingUpdate) Validate() error {
	if data.Point < 0 || data.Point > 5 {
		return ErrPointMustIn0To5
	}
	return nil
}

var ErrPointMustIn0To5 = common.NewCustomError(nil, "Point rating must be stay in 0 to 5, and is integer", "ErrPointMustIn0To5")
var ErrAlreadyRating = common.NewCustomError(nil, "User already rating restaurant", "ErrAlreadyRating")
