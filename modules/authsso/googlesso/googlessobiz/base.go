package googlessobiz

import (
	"context"
	"foodlive/modules/authsso/googlesso/googlessomodel"
	"foodlive/modules/user/usermodel"
)

type UserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type GoogleStore interface {
	ValidateGoogleJwt(ctx context.Context, tokenString string) (*googlessomodel.GoogleClaims, error)
}
