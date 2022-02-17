package ginrestaurantowner

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurantowner/restaurantownerbiz"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	"foodlive/modules/restaurantowner/restaurantownerstore"
	"github.com/gin-gonic/gin"
)

func SendOTPActiveOwnerRestaurant(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantownermodel.SendOTP

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := restaurantownerstore.NewSqlStore(appCtx.GetDatabase())

		biz := restaurantownerbiz.NewSendOTPActiveBiz(store, appCtx.GetMyCache(), appCtx.GetMySms())
		err := biz.SendOTPActiveBiz(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
