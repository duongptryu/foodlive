package ginrestaurantrating

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/restaurantrating/restaurantratingbiz"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
	"foodlive/modules/restaurantrating/restaurantratingstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateRestaurantRating(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		rstid, err := strconv.Atoi(c.Param("id_rating"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data restaurantratingmodel.RestaurantRatingUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)

		restaurantRatingStore := restaurantratingstore.NewSQLStore(appCtx.GetDatabase())
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		biz := restaurantratingbiz.UpdateRestaurantRatingBiz(restaurantStore, restaurantRatingStore, appCtx.GetPubSubProvider())

		if err := biz.UpdateRestaurantRatingBiz(c.Request.Context(), rstid, userIdRaw.(int), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
