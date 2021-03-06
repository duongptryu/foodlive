package restaurantmodel

import (
	"foodlive/common"
	"foodlive/modules/city/citymodel"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel  `json:",inline"`
	Name             string          `json:"name" gorm:"column:name;"`
	OwnerId          int             `json:"owner_id" gorm:"column:owner_id"`
	Addr             string          `json:"address" gorm:"column:addr;"`
	Logo             *common.Image   `json:"logo" gorm:"column:logo;"`
	Cover            *common.Image   `json:"cover" gorm:"column:cover;"`
	LikeCount        int             `json:"like_count" gorm:"column:liked_count"`
	IsLike           bool            `json:"is_like" gorm:"-"`
	Rating           float64         `json:"rating" gorm:"rating"`
	CityId           int             `json:"city_id" gorm:"city_id"`
	City             *citymodel.City `json:"city" gorm:"reference:CityId;preload:false"`
	Lat              float64         `json:"lat" gorm:"lat"`
	Lng              float64         `json:"lng" gorm:"lng"`
	ShippingFeePerKm float64         `json:"shipping_fee_per_km" gorm:"shipping_fee_per_km"`
	Status           bool            `json:"status" gorm:"status"`
	Distance         float64         `json:"distance" gorm:"distance"`
	RatingCount      int             `json:"rating_count" gorm:"column:rating_count"`
	TimeShipping     string          `json:"time_shipping" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name             string        `json:"name" gorm:"column:name;"`
	Addr             string        `json:"address" gorm:"column:addr;"`
	Logo             *common.Image `json:"logo" gorm:"column:logo;"`
	Cover            *common.Image `json:"cover" gorm:"column:cover;"`
	Lat              float64       `json:"lat" gorm:"lat"`
	Lng              float64       `json:"lng" gorm:"lng"`
	CityId           int           `json:"city_id" gorm:"city_id"`
	ShippingFeePerKm float64       `json:"shipping_fee_per_km" gorm:"shipping_fee_per_km"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantUpdate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpty
	}

	if res.Logo.Url == "" {
		return ErrLogoCannotBeEmpty
	}
	return nil
}

type RestaurantUpdateStatus struct {
	Status *bool `json:"status" gorm:"column:status"`
}

func (RestaurantUpdateStatus) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantUpdateStatus) Validate() error {
	if res.Status == nil {
		return ErrStatusCannotBeNil
	}
	return nil
}

type RestaurantCreate struct {
	common.SQLModel  `json:",inlines"`
	Name             string        `json:"name" gorm:"column:name;" binding:"required"`
	Addr             string        `json:"address" gorm:"column:addr;" binding:"required"`
	Logo             *common.Image `json:"logo" gorm:"column:logo;"`
	Cover            *common.Image `json:"cover" gorm:"column:cover;"`
	OwnerId          int           `json:"-" gorm:"column:owner_id"`
	Lat              float64       `json:"lat" gorm:"lat" binding:"required"`
	Lng              float64       `json:"lng" gorm:"lng" binding:"required"`
	CityId           int           `json:"city_id" gorm:"city_id" binding:"required"`
	ShippingFeePerKm float64       `json:"shipping_fee_per_km" gorm:"shipping_fee_per_km" binding:"required"`
	Status           bool          `json:"-" gorm:"status"`
	CategoryIds      []int         `json:"category_ids" gorm:"-"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpty
	}

	if res.Logo.Url == "" {
		return ErrLogoCannotBeEmpty
	}

	if len(res.CategoryIds) == 0 {
		return ErrCategoryIsRequire
	}
	return nil
}

var (
	ErrNameCannotBeEmpty  = common.NewCustomError(nil, "restaurant name can't be blank", "ErrNameCannotBeEmpty")
	ErrLogoCannotBeEmpty  = common.NewCustomError(nil, "Logo restaurant can't be empty", "ErrLogoCannotBeEmpty")
	ErrStatusCannotBeNil  = common.NewCustomError(nil, "Status restaurant can't be nil", "ErrStatusCannotBeNil")
	ErrStatusAlreadySet   = common.NewCustomError(nil, "Status restaurant already set", "ErrStatusAlreadySet")
	ErrRestaurantNotFound = common.NewFullErrorResponse(404, nil, "Restaurant not found", "Restaurant not found", "ErrRestaurantNotFound")
	ErrLatLngInvalid      = common.NewCustomError(nil, "Lat and long required", "ErrLatLngInvalid")
	ErrCategoryIsRequire  = common.NewCustomError(nil, "Category is required", "ErrCategoryIsRequire")
	ErrInvalidCategory    = common.NewCustomError(nil, "Invalid category", "ErrInvalidCategory")
)
