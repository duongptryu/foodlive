package common

type SimpleFood struct {
	SQLModel     `json:",inline"`
	RestaurantId int     `json:"restaurant_id" gorm:"column:restaurant_id"`
	CategoryId   int     `json:"category_id" gorm:"column:category_id"`
	Name         string  `json:"name" gorm:"column:name"`
	Price        float64 `json:"price" gorm:"column:price"`
}

func (SimpleFood) TableName() string {
	return "foods"
}
