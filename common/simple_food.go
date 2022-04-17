package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type SimpleFood struct {
	SQLModel     `json:",inline"`
	RestaurantId int     `json:"restaurant_id" gorm:"column:restaurant_id"`
	CategoryId   int     `json:"category_id" gorm:"column:category_id"`
	Name         string  `json:"name" gorm:"column:name"`
	Price        float64 `json:"price" gorm:"column:price"`
	Description  string  `json:"description" gorm:"column:description"`
	Images       *Image  `json:"images" gorm:"images"`
}

func (SimpleFood) TableName() string {
	return "foods"
}

type FoodOrigin struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Images *Image  `json:"images"`
}

func (j *FoodOrigin) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var fd FoodOrigin
	if err := json.Unmarshal(bytes, &fd); err != nil {
		return err
	}
	*j = fd
	return nil
}

func (j *FoodOrigin) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
