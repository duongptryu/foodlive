package ginrestaurant

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/city/citystore"
	"foodlive/modules/restaurant/restaurantbiz"
	"foodlive/modules/restaurant/restaurantmodel"
	"foodlive/modules/restaurant/restaurantstore"
	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}
		data.OwnerId = userIdRaw.(int)

		createRestaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		cityStore := citystore.NewSqlStore(appCtx.GetDatabase())
		createRestaurantBiz := restaurantbiz.NewCreateRestaurantBiz(createRestaurantStore, cityStore)

		if err := createRestaurantBiz.CreateRestaurantBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
