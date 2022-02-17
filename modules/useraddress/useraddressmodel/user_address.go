package useraddressmodel

import "foodlive/common"

const (
	EntityName = "UserAddress"
)

type UserAddress struct {
	common.SQLModel
	UserId int     `json:"user_id" gorm:"user_id"`
	CityId int     `json:"city_id" gorm:"city_id"`
	Addr   string  `json:"addr" gorm:"addr"`
	Lat    float64 `json:"lat" gorm:"lat"`
	Lng    float64 `json:"lng" gorm:"lng"`
	Status bool    `json:"status" gorm:"status"`
}

func (UserAddress) TableName() string {
	return "user_addresses"
}

type UserAddressCreate struct {
	common.SQLModelCreate
	UserId int     `json:"-" gorm:"user_id"`
	CityId int     `json:"city_id" gorm:"city_id"`
	Addr   string  `json:"addr" gorm:"addr"`
	Lat    float64 `json:"lat" gorm:"lat"`
	Lng    float64 `json:"lng" gorm:"lng"`
	Status bool    `json:"-" gorm:"status"`
}

func (UserAddressCreate) TableName() string {
	return UserAddress{}.TableName()
}

func (data *UserAddressCreate) Validate() error {
	data.Status = true
	return nil
}

type UserAddressUpdate struct {
	common.SQLModelUpdate
	CityId int     `json:"city_id" gorm:"city_id"`
	Addr   string  `json:"addr" gorm:"addr"`
	Lat    float64 `json:"lat" gorm:"lat"`
	Lng    float64 `json:"lng" gorm:"lng"`
}

func (UserAddressUpdate) TableName() string {
	return UserAddress{}.TableName()
}

func (data *UserAddressUpdate) Validate() error {
	return nil
}
