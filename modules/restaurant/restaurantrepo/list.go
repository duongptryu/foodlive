package restaurantrepo

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/restaurant/restaurantmodel"
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
//type LikeStore interface {
//	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
//}

type listRestaurantRepo struct {
	store ListRestaurantStore
	//likeStore LikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store: store,
	}
}

func (repo *listRestaurantRepo) ListRestaurantRepo(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListRestaurant(ctx, map[string]interface{}{"status": true}, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	//ids := make([]int, len(result))

	//for i := range result {
	//	ids[i] = result[i].Id
	//}

	//mapResLike, err := repo.likeStore.GetRestaurantLike(ctx, ids)
	//if err != nil {
	//	log.Println("Cannot get restaurant likes: ", err)
	//}
	//
	//if v := mapResLike; v != nil {
	//	for i, item := range result {
	//		result[i].LikeCount = mapResLike[item.Id]
	//	}
	//}

	return result, nil
}

func (repo *listRestaurantRepo) ListRestaurantOwnerRepo(ctx context.Context, userId int, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListRestaurant(ctx, map[string]interface{}{"owner_id": userId}, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	//ids := make([]int, len(result))

	//for i := range result {
	//	ids[i] = result[i].Id
	//}

	//mapResLike, err := repo.likeStore.GetRestaurantLike(ctx, ids)
	//if err != nil {
	//	log.Println("Cannot get restaurant likes: ", err)
	//}
	//
	//if v := mapResLike; v != nil {
	//	for i, item := range result {
	//		result[i].LikeCount = mapResLike[item.Id]
	//	}
	//}

	return result, nil
}
