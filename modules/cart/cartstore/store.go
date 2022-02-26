package cartstore

import (
	"context"
	"foodlive/modules/cart/cartmodel"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type CartStore interface {
	CreateCartItem(ctx context.Context, data *cartmodel.CartItemCreate) error
	DeleteCartItem(ctx context.Context, conditions map[string]interface{}) error
	FindCartItem(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*cartmodel.CartItem, error)
	UpdateCartItem(ctx context.Context, data *cartmodel.CartItemUpdate) error
	ListCartItem(ctx context.Context, condition map[string]interface{}, filter *cartmodel.Filter, moreKey ...string) ([]cartmodel.CartItem, error)
}
