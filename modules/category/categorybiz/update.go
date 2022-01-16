package categorybiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/category/categorymodel"
)

type updateCategoryBiz struct {
	categoryStore CategoryStore
}

func NewUpdateCategoryBiz(categoryStore CategoryStore) *updateCategoryBiz {
	return &updateCategoryBiz{
		categoryStore: categoryStore,
	}
}

func (biz *updateCategoryBiz) UpdateCategoryBiz(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	//check if restaurant exist
	foodDb, err := biz.categoryStore.FindCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if foodDb.Id == 0 {
		return common.ErrDataNotFound(categorymodel.EntityName)
	}

	if err := biz.categoryStore.UpdateCategory(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	return nil
}
