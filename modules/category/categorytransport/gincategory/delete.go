package gincategory

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/category/categorybiz"
	"fooddelivery/modules/category/categorystore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		categoryStore := categorystore.NewSqlStore(appCtx.GetDatabase())
		biz := categorybiz.NewDeleteCategoryBiz(categoryStore)

		if err := biz.DeleteCategoryBiz(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
