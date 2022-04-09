package ginrstcategory

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurantcategory/rstcategorybiz"
	"foodlive/modules/restaurantcategory/rstcategorymodel"
	"foodlive/modules/restaurantcategory/rstcategorystore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurantByCategory(appCtx component.AppContext) gin.HandlerFunc {
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
		biz := rstcategorybiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurantByCategory(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
