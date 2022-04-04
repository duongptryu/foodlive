package restaurantownermodel

import (
	"errors"
	"foodlive/common"
	"foodlive/component/tokenprovider"
	"time"
)

const (
	EntityName            = "OwnerRestaurant"
	TimeExpireOTPActivate = time.Second * 60
)

type OwnerRestaurant struct {
	common.SQLModel `json:",inline"`
	Phone           string `json:"phone" gorm:"column:phone"`
	Password        string `json:"-" gorm:"column:password"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:last_name"`
	Status          bool   `json:"status" gorm:"column:status"`
	Role            string `json:"role" gorm:"column:role"`
}

func (OwnerRestaurant) TableName() string {
	return "owners_restaurant"
}

type OwnerRestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Phone           string `json:"phone" gorm:"column:phone" binding:"required"`
	Password        string `json:"password" gorm:"column:password" binding:"required"`
	LastName        string `json:"last_name" gorm:"column:last_name" binding:"required"`
	FirstName       string `json:"first_name" gorm:"column:first_name" binding:"required"`
	Status          bool   `json:"-" gorm:"column:status"`
	Role            string `json:"-" gorm:"column:role"`
}

func (OwnerRestaurantCreate) TableName() string {
	return OwnerRestaurant{}.TableName()
}

type OwnerRestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Status          *bool `json:"status" gorm:"column:status"`
}

func (OwnerRestaurantUpdate) TableName() string {
	return OwnerRestaurant{}.TableName()
}

func (u *OwnerRestaurantCreate) Validate() error {
	phone := common.RePhone.Find([]byte(u.Phone))
	if phone == nil {
		return ErrPhoneInvalid
	}
	u.Phone = string(phone)

	if len(u.Password) < 8 {
		return ErrLengthPassword
	}

	if len(u.FirstName) < 3 {
		return ErrLengthFirstName
	}

	if len(u.LastName) < 3 {
		return ErrLengthLastName
	}
	u.Status = false
	u.Role = "owner_restaurant"
	return nil
}

type UserActive struct {
	Phone string `json:"phone" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

type UserLogin struct {
	Phone    string `json:"phone" gorm:"column:phone" binding:"required"`
	Password string `json:"password" gorm:"password" binding:"required"`
}

func (u *UserLogin) Validate() error {
	check := common.RePhone.Match([]byte(u.Phone))
	if !check {
		return ErrPhoneInvalid
	}

	if len(u.Password) < 8 {
		return ErrLengthPassword
	}
	return nil
}

type SendOTP struct {
	Phone string `json:"phone" gorm:"column:phone" binding:"required"`
}

func (u *SendOTP) Validate() error {
	phone := common.RePhone.Find([]byte(u.Phone))
	if phone == nil {
		return ErrPhoneInvalid
	}
	u.Phone = string(phone)
	return nil
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrPhoneInvalid             = common.NewCustomError(nil, "Phone number invalid", "ErrPhoneInvalid")
	ErrPhoneNumberAlreadyExist  = common.NewFullErrorResponse(409, errors.New("Phone number already exist"), "Phone number already exist", "Phone number already exist", "ErrPhoneNumberAlreadyExist")
	ErrPhoneNumberNotActivated  = common.NewFullErrorResponse(409, errors.New("Phone number was not activated"), "Phone number was not activated", "Phone number was not activated", "ErrPhoneNumberNotActivated")
	ErrOTPInvalidOrExpire       = common.NewCustomError(nil, "OTP invalid or expire", "ErrOTPInvalidOrExpire")
	ErrLengthPassword           = common.NewCustomError(nil, "Length of password must be greater than 8 character", "ErrLengthPassword")
	ErrLengthFirstName          = common.NewCustomError(nil, "Length of first name must be greater than 3 character", "ErrLengthFirstName")
	ErrLengthLastName           = common.NewCustomError(nil, "Length of last name must be greater than 3 character", "ErrLengthLastName")
	ErUsernameOrPasswordInvalid = common.NewFullErrorResponse(401,
		errors.New("username or password invalid"),
		"username or password invalid",
		"username or password invalid",
		"ErUsernameOrPasswordInvalid",
	)
)
