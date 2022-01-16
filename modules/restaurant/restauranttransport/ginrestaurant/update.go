package ginrestaurant

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantbiz"
	"foodlive/modules/restaurant/restaurantmodel"
	"foodlive/modules/restaurant/restaurantstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurantBiz(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}

func UpdateRestaurantStatus(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data restaurantmodel.RestaurantUpdateStatus

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurantStatusBiz(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
