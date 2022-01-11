package restaurantbiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/restaurant/restaurantmodel"
)

type ListRestaurantRepo interface {
	ListRestaurantRepo(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error)
	ListRestaurantOwnerRepo(ctx context.Context, userId int, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context, filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.repo.ListRestaurantRepo(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (biz *listRestaurantBiz) ListRestaurantOwnerBiz(ctx context.Context, userId int, filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.repo.ListRestaurantOwnerRepo(ctx, userId, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
