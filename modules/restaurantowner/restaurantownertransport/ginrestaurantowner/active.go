package ginrestaurantowner

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurantowner/restaurantownerbiz"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	"foodlive/modules/restaurantowner/restaurantownerstore"
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
