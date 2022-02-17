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
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		var filter categorymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

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
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
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
