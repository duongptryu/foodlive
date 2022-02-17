package categorybiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/category/categorymodel"
)

type deleteCategoryBiz struct {
	categoryStore CategoryStore
}

func NewDeleteCategoryBiz(categoryStore CategoryStore) *deleteCategoryBiz {
	return &deleteCategoryBiz{
		categoryStore: categoryStore,
	}
}

func (biz *deleteCategoryBiz) DeleteCategoryBiz(ctx context.Context, id int) error {
	//check if restaurant exist
	foodDb, err := biz.categoryStore.FindCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if foodDb.Id == 0 {
		return common.ErrDataNotFound(categorymodel.EntityName)
	}

	if err := biz.categoryStore.DeleteCategory(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	return nil
}
