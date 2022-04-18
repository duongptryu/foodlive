package usermodel

import (
	"errors"
	"foodlive/common"
	"foodlive/component/tokenprovider"
	"time"
)

const (
	EntityName            = "User"
	TimeExpireOTPActivate = time.Second * 60
)

type User struct {
	common.SQLModel `json:",inline"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Password        string        `json:"-" gorm:"column:password"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	FirstName       string        `json:"first_name" gorm:"column:last_name"`
	Status          bool          `json:"status" gorm:"column:status"`
	Role            string        `json:"role" gorm:"column:role"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Phone           string        `json:"phone" gorm:"column:phone" binding:"required"`
	Password        string        `json:"password" gorm:"column:password" binding:"required"`
	LastName        string        `json:"last_name" gorm:"column:last_name" binding:"required"`
	FirstName       string        `json:"first_name" gorm:"column:first_name" binding:"required"`
	Status          bool          `json:"-" gorm:"column:status"`
	Role            string        `json:"-" gorm:"column:role"`
	GgId            *string       `json:"-" gorm:"gg_id"`
	FbId            *string       `json:"-" gorm:"fb_id"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserUpdate struct {
	common.SQLModel `json:",inline"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	FirstName       string        `json:"first_name" gorm:"column:first_name"`
	Status          *bool         `json:"status" gorm:"column:status"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (data *UserUpdate) Validate() error {
	check := common.RePhone.Match([]byte(data.Phone))
	if !check {
		return ErrPhoneInvalid
	}
	return nil
}

func (u *UserCreate) Validate() error {
	check := common.RePhone.Match([]byte(u.Phone))
	if !check {
		return ErrPhoneInvalid
	}

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

type UserActive struct {
	Phone string `json:"phone" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

type UserLogin struct {
	Phone    string `json:"phone" gorm:"column:phone" binding:"required"`
	Password string `json:"password" gorm:"password" binding:"required"`
}

func (u *UserLogin) Validate() error {
	phone := common.RePhone.Find([]byte(u.Phone))
	if phone == nil {
		return ErrPhoneInvalid
	}
	u.Phone = string(phone)

	if len(u.Password) < 8 {
		return ErrLengthPassword
	}
	return nil
}

type ResendOTP struct {
	Phone string `json:"phone" gorm:"column:phone" binding:"required"`
}

func (u *ResendOTP) Validate() error {
	phone := common.RePhone.Find([]byte(u.Phone))
	if phone == nil {
		return ErrPhoneInvalid
	}
	u.Phone = string(phone)
	return nil
}

type UserResetPassword struct {
	Phone    string `json:"phone" gorm:"column:phone" binding:"required"`
	OTP      string `json:"otp" gorm:"-"`
	Password string `json:"password" gorm:"password" binding:"required"`
}

func (u *UserResetPassword) Validate() error {
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
	ErrPhoneInvalid                = common.NewCustomError(nil, "Phone number invalid", "ErrPhoneInvalid")
	ErrPhoneNumberAlreadyExist     = common.NewFullErrorResponse(409, errors.New("Phone number already exist"), "Phone number already exist", "Phone number already exist", "ErrPhoneNumberAlreadyExist")
	ErrPhoneNumberNotActivated     = common.NewFullErrorResponse(409, errors.New("Phone number was not activated"), "Phone number was not activated", "Phone number was not activated", "ErrPhoneNumberNotActivated")
	ErrPhoneNumberAlreadyActivated = common.NewFullErrorResponse(409, errors.New("Phone number already activated"), "Phone number already activated", "Phone number already activated", "ErrPhoneNumberAlreadyActivated")
	ErrOTPInvalidOrExpire          = common.NewCustomError(nil, "OTP invalid or expire", "ErrOTPInvalidOrExpire")
	ErrSendOTPMultiple             = common.NewCustomError(nil, "Please wait 60s to send the next otp code", "ErrSendOTPMultiple")
	ErrLengthPassword              = common.NewCustomError(nil, "Length of password must be greater than 8 character", "ErrLengthPassword")
	ErrLengthFirstName             = common.NewCustomError(nil, "Length of first name must be greater than 3 character", "ErrLengthFirstName")
	ErrLengthLastName              = common.NewCustomError(nil, "Length of last name must be greater than 3 character", "ErrLengthLastName")
	ErUsernameOrPasswordInvalid    = common.NewFullErrorResponse(401,
		errors.New("username or password invalid"),
		"username or password invalid",
		"username or password invalid",
		"ErUsernameOrPasswordInvalid",
	)
)
