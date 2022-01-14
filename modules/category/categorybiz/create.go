package categorybiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/category/categorymodel"
)

type createCategoryBiz struct {
	categoryStore       CategoryStore
}

func NewCreateCategoryBiz(categoryStore       CategoryStore) *createCategoryBiz {
	return &createCategoryBiz{
		categoryStore: categoryStore,
	}
}

func (biz *createCategoryBiz) CreateCategoryBiz(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	//check if restaurant exist
	result, err := biz.categoryStore.FindCategory(ctx, map[string]interface{}{"name": data.Name})
	if err != nil {
		return err
	}

	if result.Id != 0 {
		return common.ErrDataAlreadyExist(categorymodel.EntityName, "Name")
	}

	if err := biz.categoryStore.CreateCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
