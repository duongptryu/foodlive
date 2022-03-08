package userdevicetokenmodel

import (
	"foodlive/common"
	"foodlive/modules/user/usermodel"
)

const (
	EntityName = "UserDeviceToken"
)

type UserDeviceToken struct {
	common.SQLModel
	UserId int            `json:"user_id" gorm:"user_id"`
	User   usermodel.User `json:"user" gorm:"preload:false"`
	Os     string         `json:"os" gorm:"os"`
	Token  string         `json:"token" gorm:"token"`
	Status bool           `json:"status" gorm:"status"`
}

func (UserDeviceToken) TableName() string {
	return "user_device_tokens"
}

type UserDeviceTokenCreate struct {
	common.SQLModelCreate
	UserId int    `json:"-" gorm:"user_id"`
	Os     string `json:"os" gorm:"os"`
	Token  string `json:"token" gorm:"token"`
	Status bool   `json:"-" gorm:"status"`
}

func (UserDeviceTokenCreate) TableName() string {
	return UserDeviceToken{}.TableName()
}

type UserDeviceTokenUpdate struct {
	common.SQLModelUpdate
	Os    string `json:"os" gorm:"os"`
	Token string `json:"token" gorm:"token"`
}

func (UserDeviceTokenUpdate) TableName() string {
	return UserDeviceToken{}.TableName()
}
