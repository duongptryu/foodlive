package foodlikebiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/foodlike/foodlikemodel"
	"foodlive/modules/foodlike/foodlikestore"
)

type listUserLikeFood struct {
	store foodlikestore.FoodLikeStore
}

func NewlistUserLikeFood(store foodlikestore.FoodLikeStore) *listUserLikeFood {
	return &listUserLikeFood{
		store: store,
	}
}

func (biz *listUserLikeFood) ListUserLikeFood(ctx context.Context, filter *foodlikemodel.Filter,
	paging *common.Paging) ([]foodlikemodel.FoodLike, error) {
	result, err := biz.store.ListUsersLikFood(ctx, map[string]interface{}{"food_id": filter.FoodId}, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	return result, nil
}
