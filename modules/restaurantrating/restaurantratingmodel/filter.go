package restaurantratingmodel

type Filter struct {
	UserId int
	RstId  int `json:"restaurant_id" form:"restaurant_id"`
}
