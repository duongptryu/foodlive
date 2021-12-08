package usermodel

import (
	"errors"
	"fooddelivery/common"
	"fooddelivery/component/tokenprovider"
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
	Password        string `json:"-" gorm:"column:password"`
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
	Phone           string `json:"phone" gorm:"column:phone" binding:"required"`
	Password        string `json:"password" gorm:"column:password" binding:"required"`
	LastName        string `json:"last_name" gorm:"column:last_name" binding:"required"`
	FirstName       string `json:"first_name" gorm:"column:first_name" binding:"required"`
	Status          bool   `json:"-" gorm:"column:status"`
	Role            string `json:"-" gorm:"column:role"`
	GgId            string `json:"-" gorm:"gg_id"`
	FbId            string `json:"-" gorm:"fb_id"`
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
	Phone string `json:"phone" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
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

type UserLogin struct {
	Phone    string `json:"phone" gorm:"column:phone" binding:"required"`
	Password string `json:"password" gorm:"password" binding:"required"`
}

func (u *UserLogin) Validate() error {
	phone := re.Find([]byte(u.Phone))
	if phone == nil {
		return ErrPhoneInvalid
	}
	u.Phone = string(phone)

	if len(u.Password) < 8 {
		return ErrLengthPassword
	}
	return nil
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
