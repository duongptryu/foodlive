package restaurantrepo

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

type FindRestaurantStore interface {
	FindRestaurant(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

//
//type LikeStore interface {
//	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
//}

type findRestaurantRepo struct {
	store FindRestaurantStore
	//likeStore LikeStore
}

func NewFindRestaurantRepo(store FindRestaurantStore) *findRestaurantRepo {
	return &findRestaurantRepo{
		store: store,
	}
}

func (repo *findRestaurantRepo) FindRestaurantByIdRepo(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	result, err := repo.store.FindRestaurant(ctx, map[string]interface{}{"id": id, "status": true})

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	if result.Id == 0 {
		return nil, restaurantmodel.ErrRestaurantNotFound
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
