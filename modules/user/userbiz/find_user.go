package userbiz

import (
	"context"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/user/userstorage"
	"foodlive/modules/useraddress/useraddressstore"
	"foodlive/modules/userdevicetoken/userdevicetokenstore"
)

type adminFindUserBiz struct {
	userStore        userstorage.UserStore
	userAddressStore useraddressstore.UserAddressStore
	userDeviceStore  userdevicetokenstore.UserDeviceTokenStore
	orderStore       orderstore.OrderStore
}

func NewAdminFindUserBiz(userStore userstorage.UserStore, userAddressStore useraddressstore.UserAddressStore, userDeviceStore userdevicetokenstore.UserDeviceTokenStore, orderStore orderstore.OrderStore) *adminFindUserBiz {
	return &adminFindUserBiz{
		userStore:        userStore,
		userAddressStore: userAddressStore,
		userDeviceStore:  userDeviceStore,
		orderStore:       orderStore,
	}
}

func (biz *adminFindUserBiz) AdminFindUserBiz(ctx context.Context, userId int) (interface{}, error) {
	user, err := biz.userStore.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, err
	}

	userAddress, err := biz.userAddressStore.ListUserAddressByUserId(ctx, map[string]interface{}{"user_id": userId}, "City")
	if err != nil {
		return nil, err
	}

	userDevices, err := biz.userDeviceStore.FindUserDeviceToken(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}

	count, err := biz.orderStore.CountOrder(ctx, map[string]interface{}{"user_id": userId}, nil)

	result := map[string]interface{}{
		"user":        user,
		"address":     userAddress,
		"device":      userDevices,
		"order_count": count,
	}

	return &result, nil
}
