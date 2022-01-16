package restaurantownerbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	"golang.org/x/crypto/bcrypt"
)

type RegisterOwnerRestaurantStore interface {
	CreateOwnerRestaurant(ctx context.Context, data *restaurantownermodel.OwnerRestaurantCreate) error
	FindOwnerRestaurant(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*restaurantownermodel.OwnerRestaurant, error)
}

type registerOwnerRestaurantBiz struct {
	store RegisterOwnerRestaurantStore
}

func NewOwnerRestaurantBiz(store RegisterOwnerRestaurantStore) *registerOwnerRestaurantBiz {
	return &registerOwnerRestaurantBiz{
		store: store,
	}
}

func (biz *registerOwnerRestaurantBiz) RegisterOwnerRestaurantBiz(ctx context.Context, data *restaurantownermodel.OwnerRestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	userDB, err := biz.store.FindOwnerRestaurant(ctx, map[string]interface{}{"phone": data.Phone})
	if err != nil {
		return common.ErrCannotCreateEntity(restaurantownermodel.EntityName, err)
	}
	if userDB.Id != 0 {
		if userDB.Status == false {
			return restaurantownermodel.ErrPhoneNumberNotActivated
		}

		return restaurantownermodel.ErrPhoneNumberAlreadyExist
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.ErrCannotCreateEntity(restaurantownermodel.EntityName, err)
	}
	data.Password = string(hashedPassword)

	//create user in db
	if err := biz.store.CreateOwnerRestaurant(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantownermodel.EntityName, err)
	}

	return nil
}
