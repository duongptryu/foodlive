package restaurantrepo

import (
	"context"
	"fmt"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
	"log"
)

type ListRestaurantStore interface {
	ListRestaurant(ctx context.Context,
		condition map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]restaurantmodel.Restaurant, error)
}

//
type LikeStore interface {
	GetRestaurantLiked(ctx context.Context, ids []int, userId int) (map[int]bool, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore LikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store:     store,
		likeStore: likeStore,
	}
}

func (repo *listRestaurantRepo) ListRestaurantRepo(ctx context.Context, userId int, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListRestaurant(ctx, map[string]interface{}{"status": true}, filter, paging, "City")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := repo.likeStore.GetRestaurantLiked(ctx, ids, userId)
	if err != nil {
		log.Println("Cannot get restaurant liked: ", err)
	}

	if v := mapResLike; v != nil {
		for i, item := range result {
			if v, exist := mapResLike[item.Id]; exist {
				result[i].IsLike = v
			}

			//calculate time_shipping
			if item.Distance != 0 {
				result[i].TimeShipping = fmt.Sprintf("%v - %v", int(item.Distance*3), int(item.Distance*5))
			}
		}
	}

	return result, nil
}

func (repo *listRestaurantRepo) ListRestaurantOwnerRepo(ctx context.Context, userId int, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListRestaurant(ctx, map[string]interface{}{"owner_id": userId}, filter, paging, "City")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}

func (repo *listRestaurantRepo) ListRestaurantForAdmin(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListRestaurant(ctx, nil, filter, paging, "City")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
