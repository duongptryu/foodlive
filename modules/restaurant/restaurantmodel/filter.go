package restaurantmodel

type Filter struct {
	CityId  int `json:"city_id,omitempty" form:"city_id"`
	OwnerId int `json:"owner_id" form:"owner_id"`
}
