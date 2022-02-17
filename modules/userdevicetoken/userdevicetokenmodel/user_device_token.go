package userdevicetokenmodel

import "foodlive/common"

const (
	EntityName = "UserDeviceToken"
)

type UserDeviceToken struct {
	common.SQLModel
	UserId int    `json:"user_id" gorm:"user_id"`
	Os     string `json:"os" gorm:"os"`
	Token  string `json:"token" gorm:"token"`
	Status bool   `json:"status" gorm:"status"`
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
