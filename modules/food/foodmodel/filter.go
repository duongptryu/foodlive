package foodmodel

type Filter struct {
	CategoryId   int    `json:"-" form:"category"`
	RestaurantId int    `json:"-" gorm:"restaurant_id"`
	Name         string `json:"-" form:"name"`
}