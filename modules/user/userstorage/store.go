package userstorage

import (
	"context"
	"foodlive/modules/user/usermodel"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type UserStore interface {
	UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error
}
