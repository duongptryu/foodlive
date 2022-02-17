package citymodel

import "foodlive/common"

const (
	EntityName = "City"
)

type City struct {
	common.SQLModel
	Title  string `json:"title" gorm:"title"`
	Status bool   `json:"status" gorm:"status"`
}

func (City) TableName() string {
	return "cities"
}

type CityUpdate struct {
	common.SQLModelUpdate
	Status *bool `json:"status" gorm:"status"`
}

func (CityUpdate) TableName() string {
	return City{}.TableName()
}
