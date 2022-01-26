package ginuserdevicetoken

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/userdevicetoken/userdevicetokenbiz"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
	"foodlive/modules/userdevicetoken/userdevicetokenstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUserDeviceToken(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter userdevicetokenmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		store := userdevicetokenstore.NewSQLStore(appCtx.GetDatabase())
		biz := userdevicetokenbiz.NewListUserDeviceTokenBiz(store)

		result, err := biz.ListUserDeviceTokenBiz(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
