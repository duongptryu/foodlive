package rstcategorybiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantcategory/rstcategorymodel"
	"foodlive/modules/restaurantcategory/rstcategorystore"
)

type listRstCategoryBiz struct {
	store rstcategorystore.RstCategoryStore
}

func NewListRestaurantBiz(store rstcategorystore.RstCategoryStore) *listRstCategoryBiz {
	return &listRstCategoryBiz{
		store: store,
	}
}

func (biz *listRstCategoryBiz) ListRestaurantByCategory(ctx context.Context, filter *rstcategorymodel.Filter,
	paging *common.Paging) ([]rstcategorymodel.RstCategory, error) {

	result, err := biz.store.ListRestaurantByCategory(ctx, map[string]interface{}{"category_id": filter.CategoryId}, filter, paging, "Restaurant")
	if err != nil {
		return nil, err
	}
	return result, nil
}
