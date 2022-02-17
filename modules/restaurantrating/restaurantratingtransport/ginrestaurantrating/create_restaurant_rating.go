package ginrestaurantrating

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/restaurantrating/restaurantratingbiz"
	"foodlive/modules/restaurantrating/restaurantratingmodel"
	"foodlive/modules/restaurantrating/restaurantratingstore"
	"github.com/gin-gonic/gin"
)

func CreateRestaurantRating(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantratingmodel.RestaurantRatingCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		data.UserId = userIdRaw.(int)

		restaurantRatingStore := restaurantratingstore.NewSQLStore(appCtx.GetDatabase())
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		biz := restaurantratingbiz.NewCreateRestaurantRatingBiz(restaurantStore, restaurantRatingStore, appCtx.GetPubSubProvider())

		if err := biz.CreateRestaurantRatingBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
