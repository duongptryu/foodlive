package restaurantownermodel

type Filter struct {
	Phone    string `json:"phone" form:"phone"`
	LastName string `json:"last_name" form:"last_name"`
}
