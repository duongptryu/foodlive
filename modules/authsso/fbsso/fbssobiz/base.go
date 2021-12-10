package fbssobiz

import (
	"context"
	"fooddelivery/modules/authsso/fbsso/fbssomodel"
	"fooddelivery/modules/user/usermodel"
)

type UserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type FbStore interface {
	ValidateFbJwt(ctx context.Context, tokenString string) (*fbssomodel.FacebookUser, error)
}
