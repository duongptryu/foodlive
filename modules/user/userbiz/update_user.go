package userbiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/user/usermodel"
	"foodlive/modules/user/userstorage"
)

type updateUserBiz struct {
	store userstorage.UserStore
}

func NewUpdateUserBiz(store userstorage.UserStore) *updateUserBiz {
	return &updateUserBiz{
		store: store,
	}
}

func (biz *updateUserBiz) UpdateUserBiz(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	if err := biz.store.UpdateUser(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
