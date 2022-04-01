package statisticbiz

import (
	"context"
	"foodlive/component/paymentprovider"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/statistic/statisticmodel"
	"foodlive/modules/user/userstorage"
	"time"
)

type overviewBiz struct {
	orderStore      orderstore.OrderStore
	userStore       userstorage.UserStore
	rstStore        restaurantstore.RestaurantStore
	rinkebyProvider paymentprovider.CryptoPaymentProvider
}

func NewOverviewBiz(orderStore orderstore.OrderStore, userStore userstorage.UserStore, rstStore restaurantstore.RestaurantStore, rinkebyProvider paymentprovider.CryptoPaymentProvider) *overviewBiz {
	return &overviewBiz{
		orderStore:      orderStore,
		userStore:       userStore,
		rstStore:        rstStore,
		rinkebyProvider: rinkebyProvider,
	}
}

func (biz *overviewBiz) OverviewBiz(ctx context.Context) (*statisticmodel.OverviewResp, error) {
	now := time.Now()
	todayCondition := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	totalUser, err := biz.userStore.CountUser(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	todayUser, err := biz.userStore.CountUser(ctx, nil, &todayCondition)
	if err != nil {
		return nil, err
	}

	totalRst, err := biz.rstStore.CountRst(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	todayRst, err := biz.rstStore.CountRst(ctx, nil, &todayCondition)
	if err != nil {
		return nil, err
	}

	orders, err := biz.orderStore.ListOrderWithoutPaging(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	todayOrder, err := biz.orderStore.CountOrder(ctx, nil, &todayCondition)
	if err != nil {
		return nil, err
	}

	var totalMoney float64

	for i, _ := range orders {
		totalMoney += orders[i].TotalPrice
	}

	result := statisticmodel.OverviewResp{
		totalUser,
		todayUser,
		len(orders),
		todayOrder,
		totalRst,
		todayRst,
		totalMoney,
	}

	return &result, nil
}
