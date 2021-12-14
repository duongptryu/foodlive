package restaurantmodel

import (
	"fooddelivery/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	OwnerId         int            `json:"-" gorm:"column:owner_id"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
	//User            *common.SimpleUser `json:"user" gorm:"preload:false"`
	//LikeCount int `json:"like_count" gorm:"column:liked_count"`
	CityId int `json:"city_id" gorm:"city_id"`
	//City   *citymodel.City `json:"city" gorm:"preload:false"`
	Lat              float64 `json:"lat" gorm:"lat"`
	Lng              float64 `json:"lng" gorm:"lng"`
	ShippingFeePerKm float64 `json:"shipping_fee_per_km" gorm:"shipping_fee_per_km"`
	Status           bool    `json:"status" gorm:"status"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name             string         `json:"name" gorm:"column:name;"`
	Addr             string         `json:"address" gorm:"column:addr;"`
	Logo             *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover            *common.Images `json:"cover" gorm:"column:cover;"`
	Lat              float64        `json:"lat" gorm:"lat"`
	Lng              float64        `json:"lng" gorm:"lng"`
	CityId           int            `json:"city_id" gorm:"city_id"`
	ShippingFeePerKm float64        `json:"shipping_fee_per_km" gorm:"shipping_fee_per_km"`
	Status           bool           `json:"status" gorm:"status"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel  `json:",inlines"`
	Name             string         `json:"name" gorm:"column:name;" binding:"required"`
	Addr             string         `json:"address" gorm:"column:addr;" binding:"required"`
	Logo             *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover            *common.Images `json:"cover" gorm:"column:cover;"`
	OwnerId          int            `json:"-" gorm:"column:owner_id"`
	Lat              float64        `json:"lat" gorm:"lat" binding:"required"`
	Lng              float64        `json:"lng" gorm:"lng" binding:"required"`
	CityId           int            `json:"city_id" gorm:"city_id" binding:"required"`
	ShippingFeePerKm float64        `json:"-" gorm:"shipping_fee_per_km"`
	Status           bool           `json:"-" gorm:"status"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpty
	}

	return nil
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "restaurant name can't be blank", "ErrNameCannotBeEmpty")
)
