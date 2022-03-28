package ginrestaurantowner

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurantowner/restaurantownerbiz"
	"foodlive/modules/restaurantowner/restaurantownermodel"
	"foodlive/modules/restaurantowner/restaurantownerstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AdminUpdateOwnerRst(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		ownerId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data restaurantownermodel.OwnerRestaurantUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		ownerStore := restaurantownerstore.NewSqlStore(appCtx.GetDatabase())
		userActiveBiz := restaurantownerbiz.NewUpdateOwnerRstBiz(ownerStore)

		if err := userActiveBiz.UpdateOwnerRstBiz(c.Request.Context(), ownerId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
