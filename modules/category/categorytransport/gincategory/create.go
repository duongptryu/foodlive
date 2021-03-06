package gincategory

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/category/categorybiz"
	"foodlive/modules/category/categorymodel"
	"foodlive/modules/category/categorystore"
	"github.com/gin-gonic/gin"
)

func CreateCategory(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		categoryStore := categorystore.NewSqlStore(appCtx.GetDatabase())
		biz := categorybiz.NewCreateCategoryBiz(categoryStore)

		if err := biz.CreateCategoryBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
