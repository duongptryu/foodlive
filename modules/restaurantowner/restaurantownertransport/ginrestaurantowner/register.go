package ginrestaurantowner

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurantowner/restaurantownerbiz"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	"foodlive/modules/restaurantowner/restaurantownerstore"
	"github.com/gin-gonic/gin"
)

func OwnerRestaurantRegister(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantownermodel.OwnerRestaurantCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userRegisterStore := restaurantownerstore.NewSqlStore(appCtx.GetDatabase())
		userRegisterBiz := restaurantownerbiz.NewOwnerRestaurantBiz(userRegisterStore)

		if err := userRegisterBiz.RegisterOwnerRestaurantBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
