package useraddressmodel

import (
	"foodlive/common"
	"foodlive/modules/city/citymodel"
)

const (
	EntityName = "UserAddress"
)

type UserAddress struct {
	common.SQLModel
	UserId    int             `json:"user_id" gorm:"user_id"`
	CityId    int             `json:"city_id" gorm:"city_id"`
	Name      string          `json:"name" gorm:"name"`
	Phone     string          `json:"phone" gorm:"phone"`
	City      *citymodel.City `json:"city" gorm:"preload:false"`
	Addr      string          `json:"address" gorm:"addr"`
	Lat       float64         `json:"lat" gorm:"lat"`
	Lng       float64         `json:"lng" gorm:"lng"`
	Status    bool            `json:"status" gorm:"status"`
	IsDefault bool            `json:"is_default" gorm:"column:is_default"`
}

func (UserAddress) TableName() string {
	return "user_addresses"
}

type UserAddressCreate struct {
	common.SQLModelCreate
	UserId    int     `json:"-" gorm:"user_id"`
	Name      string  `json:"name" gorm:"name"`
	Phone     string  `json:"phone" gorm:"phone"`
	CityId    int     `json:"city_id" gorm:"city_id"`
	Addr      string  `json:"address" gorm:"addr"`
	Lat       float64 `json:"lat" gorm:"lat"`
	Lng       float64 `json:"lng" gorm:"lng"`
	Status    bool    `json:"-" gorm:"status"`
	IsDefault bool    `json:"is_default" gorm:"column:is_default"`
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
	CityId    int     `json:"city_id" gorm:"city_id"`
	Addr      string  `json:"address" gorm:"addr"`
	Name      string  `json:"name" gorm:"name"`
	Phone     string  `json:"phone" gorm:"phone"`
	Lat       float64 `json:"lat" gorm:"lat"`
	IsDefault *bool   `json:"is_default" gorm:"column:is_default"`
	Lng       float64 `json:"lng" gorm:"lng"`
}

func (UserAddressUpdate) TableName() string {
	return UserAddress{}.TableName()
}

func (data *UserAddressUpdate) Validate() error {
	return nil
}

var (
	ErrAlreadyHasDefaultAddress      = common.NewCustomError(nil, "user already has default address", "ErrAlreadyHasDefaultAddress")
	ErrUserDoseNotHaveDefaultAddress = common.NewCustomError(nil, "User does not have default address", "ErrUserDoseNotHaveDefaultAddress")
)
