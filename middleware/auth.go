package middleware

import (
	"context"
	"errors"
	"fmt"
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/user/usermodel"
	"foodlive/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extracTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequireAuth(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := appCtx.GetTokenProvider()
	//myCache := appCtx.GetMyCache()
	userStore := userstorage.NewSQLStore(appCtx.GetDatabase())
	return func(c *gin.Context) {
		token, err := extracTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		if payload.Type == common.TypeAccountSocial {
			//check if account set phone number
			user, err := userStore.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
			if err != nil {
				panic(err)
			}
			if user.Phone == "" {
				panic(common.NewFullErrorResponse(409, nil, "Account must update phone number", "", "ErrMissingPhoneNumber"))
			}
			if user.Status == false {
				panic(usermodel.ErrPhoneNumberNotActivated)
			}
		}

		c.Set(common.KeyUserHeader, payload.UserId)
		c.Next()
	}
}

func RequireSSOAuth(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := appCtx.GetTokenProvider()

	return func(c *gin.Context) {
		token, err := extracTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		if payload.Type != common.TypeAccountSocial {
			panic(common.ErrDataNotFound("Page Not Found"))
		}
		c.Set(common.KeyUserHeader, payload.UserId)
		c.Next()
	}
}

func RequireAuthAdmin(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := appCtx.GetTokenProvider()
	//myCache := appCtx.GetMyCache()
	return func(c *gin.Context) {
		token, err := extracTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		if payload.Role != "admin" {
			panic(ErrInvalidToken)
		}
		//
		//userId, err := myCache.Get(common.KeyTokenCache + token)
		//if err != nil {
		//	panic(ErrNotFound)
		//}
		//
		//if payload.UserId != userId.(int) {
		//	panic(ErrInvalidToken)
		//}

		c.Set(common.KeyUserHeader, payload.UserId)
		c.Next()
	}
}

func RequireAuthOwnerRestaurant(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := appCtx.GetTokenProvider()
	//myCache := appCtx.GetMyCache()
	return func(c *gin.Context) {
		token, err := extracTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		if payload.Role != "owner_restaurant" {
			panic(ErrInvalidToken)
		}

		c.Set(common.KeyUserHeader, payload.UserId)
		c.Next()
	}
}

var (
	ErrNotFound = common.NewFullErrorResponse(401,
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
		"ErrNotFound",
	)
	ErrInvalidToken = common.NewFullErrorResponse(401, errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
		"ErrInvalidToken",
	)
)
