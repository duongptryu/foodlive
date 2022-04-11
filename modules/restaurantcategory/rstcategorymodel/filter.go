package rstcategorymodel

type Filter struct {
	CategoryId   int `json:"category_id" form:"category_id"`
	RestaurantId int `json:"restaurant_id" form:"restaurant_id"`
}
