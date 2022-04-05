package userstorage

import (
	"context"
	"foodlive/modules/user/usermodel"
	"gorm.io/gorm"
	"time"
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
	CountUser(ctx context.Context, condition map[string]interface{}, conditionTime *time.Time) (int, error)
	ListUserWithoutPaging(ctx context.Context,
		condition map[string]interface{},
		filter *usermodel.Filter,
		moreKey ...string,
	) ([]usermodel.User, error)
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*usermodel.User, error)
}
