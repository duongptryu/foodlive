package ginrestaurant

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/restaurant/restaurantbiz"
	"fooddelivery/modules/restaurant/restaurantstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		store := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurantBiz(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
