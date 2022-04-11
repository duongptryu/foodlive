package ginrstcategory

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurantcategory/rstcategorybiz"
	"foodlive/modules/restaurantcategory/rstcategorymodel"
	"foodlive/modules/restaurantcategory/rstcategorystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategoryInRst(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter rstcategorymodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}

		paging.Fulfill()

		store := rstcategorystore.NewSqlStore(appCtx.GetDatabase())
		biz := rstcategorybiz.NewListCategoryInRstBiz(store)

		result, err := biz.ListCategoryInRstBiz(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
