package foodmodel

type Filter struct {
	CategoryId int    `json:"-" form:"category"`
	Name       string `json:"-" form:"name"`
}
