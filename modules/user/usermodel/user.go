package usermodel

import (
	"errors"
	"fooddelivery/common"
	"math/rand"
	"regexp"
	"time"
)

//var re = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
var re = regexp.MustCompile(`[84|]+(3|5|7|8|9|1[2|6|8|9])+([0-9]{8})`)

const (
	EntityName            = "User"
	EntityOTP             = "OTP"
	numberStr             = "0123456789"
	TimeExpireOTPActivate = time.Second * 60
)

type User struct {
	common.SQLModel `json:",inline"`
	Phone           string `json:"phone" gorm:"column:phone"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:last_name"`
	Status          bool   `json:"status" gorm:"column:status"`
	Role            string `json:"role" gorm:"column:role"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Phone           string `json:"phone" gorm:"column:phone" binding:"required,min=9,max=20"`
	Password        string `json:"password" gorm:"column:password" binding:"required,min=6"`
	LastName        string `json:"last_name" gorm:"column:last_name" binding:"required,min=3"`
	FirstName       string `json:"first_name" gorm:"column:first_name" binding:"required,min=3"`
	Status          bool   `json:"-" gorm:"column:status"`
	Role            string `json:"-" gorm:"column:role"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Validate() error {
	phone := re.Find([]byte(u.Phone))
	if phone == nil {
		return ErrPhoneInvalid
	}
	u.Phone = string(phone)

	u.Status = false
	u.Role = "user"
	return nil
}

func GenerateOTP(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numberStr[rand.Intn(len(numberStr))]
	}
	return string(b)
}

type UserActive struct {
	Phone string `json:"phone" binding:"required,min=9,max=20"`
	OTP   string `json:"otp" binding:"required,min=4,max=4"`
}

var (
	ErrPhoneInvalid            = common.NewCustomError(nil, "Phone number invalid", "ErrPhoneInvalid")
	ErrPhoneNumberAlreadyExist = common.NewFullErrorResponse(409, errors.New("Phone number already exist"), "Phone number already exist", "Phone number already exist", "ErrPhoneNumberAlreadyExist")
	ErrPhoneNumberNotActivated = common.NewFullErrorResponse(409, errors.New("Phone number was not activated"), "Phone number was not activated", "Phone number was not activated", "ErrPhoneNumberNotActivated")
	ErrOTPInvalidOrExpire      = common.NewCustomError(nil, "OTP invalid or expire", "ErrOTPInvalidOrExpire")
)
