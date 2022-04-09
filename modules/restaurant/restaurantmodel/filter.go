package restaurantmodel

type Filter struct {
	CityId     int     `json:"city_id,omitempty" form:"city_id"`
	OwnerId    int     `json:"owner_id" form:"owner_id"`
	Lat        float64 `json:"lat" form:"lat"`
	Lng        float64 `json:"lng" form:"lng"`
	Distance   float64 `json:"distance" form:"distance"`
	OrderBy    string  `json:"order_by" form:"order_by"`
	Name       string  `json:"name" form:"name"`
	Address    string  `json:"address" form:"address"`
	CategoryId int     `json:"category_id" form:"category_id"`
}
