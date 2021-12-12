package ginrestaurantowner

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/restaurantowner/restaurantownerbiz"
	"fooddelivery/modules/restaurantowner/restaurantownermodel"
	"fooddelivery/modules/restaurantowner/restaurantownerstore"
	"github.com/gin-gonic/gin"
)

func OwnerRestaurantActive(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantownermodel.UserActive

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userActiveStore := restaurantownerstore.NewSqlStore(appCtx.GetDatabase())
		userActiveBiz := restaurantownerbiz.NewActiveUserBiz(userActiveStore, appCtx.GetMyCache())

		if err := userActiveBiz.ActiveOwnerRestaurantBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
