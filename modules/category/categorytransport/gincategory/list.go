package gincategory

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/category/categorybiz"
	"foodlive/modules/category/categorymodel"
	"foodlive/modules/category/categorystore"
	"github.com/gin-gonic/gin"
)

func AdminListCategory(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var filter categorymodel.Filter
		if err := c.ShouldBindJSON(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		categoryStore := categorystore.NewSqlStore(appCtx.GetDatabase())
		biz := categorybiz.NewListCategoryBiz(categoryStore)

		result, err := biz.AdminListCategoryBiz(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}

func UserListCategory(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var filter categorymodel.Filter
		if err := c.ShouldBindJSON(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		categoryStore := categorystore.NewSqlStore(appCtx.GetDatabase())
		biz := categorybiz.NewListCategoryBiz(categoryStore)

		result, err := biz.UserListCategoryBiz(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}
