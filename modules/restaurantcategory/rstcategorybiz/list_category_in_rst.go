package rstcategorybiz

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantcategory/rstcategorymodel"
	"foodlive/modules/restaurantcategory/rstcategorystore"
)

type listCategoryInRstBiz struct {
	store rstcategorystore.RstCategoryStore
}

func NewListCategoryInRstBiz(store rstcategorystore.RstCategoryStore) *listCategoryInRstBiz {
	return &listCategoryInRstBiz{
		store: store,
	}
}

func (biz *listCategoryInRstBiz) ListCategoryInRstBiz(ctx context.Context, filter *rstcategorymodel.Filter,
	paging *common.Paging) ([]rstcategorymodel.RstCategory, error) {

	result, err := biz.store.ListRestaurantByCategory(ctx, map[string]interface{}{"category_id": filter.CategoryId}, filter, paging, "Restaurant")
	if err != nil {
		return nil, err
	}
	return result, nil
}
