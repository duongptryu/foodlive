package rstcategorymodel

import (
	"time"
)

type RstCategory struct {
	CategoryId int `json:"category_id"`
	//Category        *categorymodel.Category     `json:"category" gorm:"preload:false"`
	RestaurantId int `json:"restaurant_id" gorm:"restaurant_id"`
	//Restaurant      *restaurantmodel.Restaurant `json:"restaurant" gorm:"preload:false"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (RstCategory) TableName() string {
	return "restaurant_category"
}

type RstCategoryCreate struct {
	CategoryId   int        `json:"category_id"`
	RestaurantId int        `json:"restaurant_id" gorm:"restaurant_id"`
	CreatedAt    *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"updated_at"`
}

func (RstCategoryCreate) TableName() string {
	return RstCategory{}.TableName()
}
