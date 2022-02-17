package categorybiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/category/categorymodel"
	"foodlive/modules/food/foodmodel"
)

type listCategoryBiz struct {
	categoryStore CategoryStore
}

func NewListCategoryBiz(categoryStore CategoryStore) *listCategoryBiz {
	return &listCategoryBiz{
		categoryStore: categoryStore,
	}
}

func (biz *listCategoryBiz) AdminListCategoryBiz(ctx context.Context, paging *common.Paging, filter *categorymodel.Filter) ([]categorymodel.Category, error) {
	result, err := biz.categoryStore.ListCategory(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return result, nil
}

func (biz *listCategoryBiz) UserListCategoryBiz(ctx context.Context, paging *common.Paging, filter *categorymodel.Filter) ([]categorymodel.Category, error) {
	result, err := biz.categoryStore.ListCategory(ctx, map[string]interface{}{"status": true}, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return result, nil
}
