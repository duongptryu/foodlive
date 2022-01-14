package categorymodel

type Filter struct {
	Name string `json:"-" form:"name"`
	Status bool `json:"-" form:"status"`
}
