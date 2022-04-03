package restaurantmodel

type Filter struct {
	CityId   int     `json:"city_id,omitempty" form:"city_id"`
	OwnerId  int     `json:"owner_id" form:"owner_id"`
	Lat      float64 `json:"lat" form:"lat"`
	Lng      float64 `json:"lng" form:"lng"`
	Distance float64 `json:"distance" form:"distance"`
	OrderBy  string  `json:"order_by" form:"order_by"`
}
