package gincategory

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/category/categorybiz"
	"fooddelivery/modules/category/categorymodel"
	"fooddelivery/modules/category/categorystore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data categorymodel.CategoryUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		categoryStore := categorystore.NewSqlStore(appCtx.GetDatabase())
		biz := categorybiz.NewUpdateCategoryBiz(categoryStore)

		if err := biz.UpdateCategoryBiz(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
