package restaurantbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurant/restaurantmodel"
)

type ListRestaurantRepo interface {
	ListRestaurantRepo(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error)
	ListRestaurantOwnerRepo(ctx context.Context, userId int, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error)
	ListRestaurantForAdmin(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context, filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	if filter == nil || (filter.Lng == 0 && filter.Lat == 0) {
		return nil, restaurantmodel.ErrLatLngInvalid
	}

	result, err := biz.repo.ListRestaurantRepo(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (biz *listRestaurantBiz) ListRestaurantOwner(ctx context.Context, userId int, filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.repo.ListRestaurantOwnerRepo(ctx, userId, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (biz *listRestaurantBiz) ListRestaurantForAdmin(ctx context.Context, filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.repo.ListRestaurantForAdmin(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}