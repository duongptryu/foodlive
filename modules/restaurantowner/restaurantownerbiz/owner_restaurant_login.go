package restaurantownerbiz

import (
	"context"
	"foodlive/common"
	"foodlive/component/mycache"
	"foodlive/component/tokenprovider"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type OwnerRestaurantLoginStore interface {
	FindOwnerRestaurant(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*restaurantownermodel.OwnerRestaurant, error)
}

type LoginHashProvider interface {
	ValidatePassword(password, hasedPassword string) (bool, error)
}

type ownerRestaurantLoginBiz struct {
	storeUser     OwnerRestaurantLoginStore
	myCache       mycache.Cache
	tokenProvider tokenprovider.TokenProvider
	expire        int
}

func NewOwnerRestaurantLoginBiz(storeUser OwnerRestaurantLoginStore, myCache mycache.Cache, tokenProvider tokenprovider.TokenProvider, expiry int) *ownerRestaurantLoginBiz {
	return &ownerRestaurantLoginBiz{
		storeUser:     storeUser,
		myCache:       myCache,
		tokenProvider: tokenProvider,
		expire:        expiry,
	}
}

func (biz *ownerRestaurantLoginBiz) OwnerRestaurantLoginBiz(ctx context.Context, data *restaurantownermodel.UserLogin) (*restaurantownermodel.Account, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	userDB, err := biz.storeUser.FindOwnerRestaurant(ctx, map[string]interface{}{"phone": data.Phone})
	if err != nil {
		return nil, restaurantownermodel.ErUsernameOrPasswordInvalid
	}

	if userDB.Id == 0 {
		return nil, restaurantownermodel.ErUsernameOrPasswordInvalid
	}

	if userDB.Status == false {
		return nil, restaurantownermodel.ErrPhoneNumberNotActivated
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(data.Password))
	if err != nil {
		return nil, restaurantownermodel.ErUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: userDB.Id,
		Role:   userDB.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(&payload, biz.expire)

	if err != nil {
		return nil, common.ErrInternal(err)
	}
	//
	//refreshToken, err := biz.tokenProvider.Generate(&payload, biz.expire)
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}

	account := restaurantownermodel.NewAccount(accessToken, nil)

	err = biz.myCache.SetWithExpire(common.KeyTokenCache+accessToken.Token, userDB.Id, common.TimeExpireTokenCache)
	if err != nil {
		log.Error(err)
	}
	return account, nil
}
