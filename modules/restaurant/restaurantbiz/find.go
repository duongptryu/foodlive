package restaurantbiz

import (
	"context"
	"foodlive/modules/restaurant/restaurantmodel"
)

type FindRestaurantRepo interface {
	FindRestaurantByIdRepo(ctx context.Context, id int, userId int) (*restaurantmodel.Restaurant, error)
	FindRestaurantByIdRepoWithoutStatus(ctx context.Context, id int) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBiz struct {
	repo FindRestaurantRepo
}

func NewFindRestaurantBiz(repo FindRestaurantRepo) *findRestaurantBiz {
	return &findRestaurantBiz{repo: repo}
}

func (biz *findRestaurantBiz) FindRestaurantById(ctx context.Context, id int, userId int) (*restaurantmodel.Restaurant, error) {
	result, err := biz.repo.FindRestaurantByIdRepo(ctx, id, userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (biz *findRestaurantBiz) FindRestaurantByIdWithoutStatus(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	result, err := biz.repo.FindRestaurantByIdRepoWithoutStatus(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
