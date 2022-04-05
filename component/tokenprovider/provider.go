package tokenprovider

import (
	"errors"
	"foodlive/common"
	"time"
)

type TokenProvider interface {
	Generate(data *TokenPayload, expire int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expire  int       `json:"expire"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	Type   string `json:"type"`
}

var (
	ErrNotFound = common.NewFullErrorResponse(401,
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewFullErrorResponse(401, errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
		"ErrInvalidToken",
	)
)
