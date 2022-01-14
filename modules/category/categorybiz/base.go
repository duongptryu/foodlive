package categorybiz

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/category/categorymodel"
)

type CategoryStore interface {
	CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error
	DeleteCategory(ctx context.Context, id int) error
	FindCategory(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*categorymodel.Category, error)
	ListCategory(ctx context.Context, condition map[string]interface{}, filter *categorymodel.Filter, paging *common.Paging, moreKey ...string) ([]categorymodel.Category, error)
	UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error
}
