package googlessomodel

import (
	"foodlive/common"
	"foodlive/component/tokenprovider"
	"github.com/dgrijalva/jwt-go"
)

const EntityName = "UserGoogle"

type GoogleJwtInput struct {
	Token string `json:"token" binding:"required"`
}

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	jwt.StandardClaims
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
	ErrInvalidGoogleJwt        = common.NewFullErrorResponse(401, nil, "Google JWT invalid", "Google JWT invalid", "ErrInvalidGoogleJwt")
	ErrGoogleJwtExpire         = common.NewFullErrorResponse(401, nil, "Google JWT is expired", "Google JWT is expired", "ErrGoogleJwtExpire")
	ErrInvalidJwtIss           = common.NewFullErrorResponse(401, nil, "Jwt iss invalid", "Jwt iss invalid", "ErrInvalidJwtIss")
	ErrGoogleKeyNotFound       = common.NewFullErrorResponse(401, nil, "Google key not found", "Google key not found", "ErrGoogleKeyNotFound")
	ErrUserPhoneEmpty          = common.NewFullErrorResponse(409, nil, "User phone number is empty", "User phone number is empty", "ErrUserPhoneEmpty")
	ErrPhoneNumberNotActivated = common.NewFullErrorResponse(409, nil, "Phone number was not activated", "Phone number was not activated", "ErrPhoneNumberNotActivated")
	ErrAccountAlreadyExisted   = common.NewFullErrorResponse(409, nil, "Account already existed", "Account already existed", "ErrAccountAlreadyExisted")
	ErrAccountDoesNotExist     = common.NewFullErrorResponse(409, nil, "Account does not exist", "Account does not exist", "ErrAccountDoesNotExist")
)
