package ginrestaurantowner

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/restaurantowner/restaurantownerbiz"
	"fooddelivery/modules/restaurantowner/restaurantownermodel"
	"fooddelivery/modules/restaurantowner/restaurantownerstore"
	"github.com/gin-gonic/gin"
)

func OwnerRestaurantLogin(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantownermodel.UserLogin

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := restaurantownerstore.NewSqlStore(appCtx.GetDatabase())

		biz := restaurantownerbiz.NewOwnerRestaurantLoginBiz(store, appCtx.GetTokenProvider(), 60*60*24*30)
		account, err := biz.OwnerRestaurantLoginBiz(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(account))
	}
}
