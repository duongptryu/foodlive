package restaurantownerbiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/mycache"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	log "github.com/sirupsen/logrus"
)

type ActiveOwnerRestaurantStore interface {
	UpdateStatusOwnerRestaurant(ctx context.Context, phoneNumber string) error
}

type activateOwnerRestaurantBiz struct {
	store   ActiveOwnerRestaurantStore
	myCache mycache.Cache
}

func NewActiveUserBiz(store ActiveOwnerRestaurantStore, myCache mycache.Cache) *activateOwnerRestaurantBiz {
	return &activateOwnerRestaurantBiz{
		store:   store,
		myCache: myCache,
	}
}

func (biz *activateOwnerRestaurantBiz) ActiveOwnerRestaurantBiz(ctx context.Context, data *restaurantownermodel.UserActive) error {
	otp, err := biz.myCache.Get(common.EntityOTP + data.Phone)
	if err != nil {
		return restaurantownermodel.ErrOTPInvalidOrExpire
	}

	if otp.(string) != data.OTP {
		return restaurantownermodel.ErrOTPInvalidOrExpire
	}

	//remove cache
	if err := biz.myCache.Remove(common.EntityOTP + data.Phone); err != nil {
		log.Error("Cannot remove OTP in cache - ", err)
	}

	if err := biz.store.UpdateStatusOwnerRestaurant(ctx, data.Phone); err != nil {
		return common.ErrCannotUpdateEntity(restaurantownermodel.EntityName, err)
	}
	return nil
}
