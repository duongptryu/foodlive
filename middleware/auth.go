package middleware

import (
	"errors"
	"fmt"
	"foodlive/common"
	"foodlive/component"
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
	return func(c *gin.Context) {
		token, err := extracTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
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

		//userId, err := myCache.Get(common.KeyTokenCache + token)
		//if err != nil {
		//	panic(ErrNotFound)
		//}

		//if payload.UserId != userId.(int) {
		//	panic(ErrInvalidToken)
		//}

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
