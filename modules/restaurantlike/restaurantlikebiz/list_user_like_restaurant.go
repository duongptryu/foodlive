package restaurantlikebiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
)

type ListUserLikRestaurantStore interface {
	GetUsersLikeRestaurant(ctx context.Context, conditions map[string]interface{}, filter *restaurantlikemodel.Filter, paging *common.Paging, moreKeys ...string) ([]common.SimpleUser, error)
}

type listUserLikeRestaurant struct {
	store ListUserLikRestaurantStore
}

func NewListUserLikeRestaurant(store ListUserLikRestaurantStore) *listUserLikeRestaurant {
	return &listUserLikeRestaurant{
		store: store,
	}
}

func (biz *listUserLikeRestaurant) ListUser(ctx context.Context, filter *restaurantlikemodel.Filter,
	paging *common.Paging) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
