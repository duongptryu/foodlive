package fbssomodel

import (
	"foodlive/common"
	"foodlive/component/tokenprovider"
)

const EntityName = "UserFacebook"

type FacebookUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FacebookJwtInput struct {
	Token string `json:"token" binding:"required"`
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
	ErrInvalidFacebookJwt      = common.NewFullErrorResponse(401, nil, "Invalid facebook token", "Invalid facebook token", "ErrInvalidFacebookJwt")
	ErrUserPhoneEmpty          = common.NewFullErrorResponse(409, nil, "User phone number is empty", "User phone number is empty", "ErrUserPhoneEmpty")
	ErrPhoneNumberNotActivated = common.NewFullErrorResponse(409, nil, "Phone number was not activated", "Phone number was not activated", "ErrPhoneNumberNotActivated")
	ErrAccountAlreadyExisted   = common.NewFullErrorResponse(409, nil, "Account already existed", "Account already existed", "ErrAccountAlreadyExisted")
	ErrAccountDoesNotExist     = common.NewFullErrorResponse(409, nil, "Account does not exist", "Account does not exist", "ErrAccountDoesNotExist")
)
