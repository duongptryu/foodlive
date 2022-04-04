package ginrestaurantrating

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/restaurantrating/restaurantratingbiz"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
	"foodlive/modules/restaurantrating/restaurantratingstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListMyRating(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantratingmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userId := c.MustGet(common.KeyUserHeader)
		filter.UserId = userId.(int)

		restaurantRatingStore := restaurantratingstore.NewSQLStore(appCtx.GetDatabase())
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		biz := restaurantratingbiz.NewListMyRatingBiz(restaurantStore, restaurantRatingStore)

		result, err := biz.ListMyRatingBiz(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
