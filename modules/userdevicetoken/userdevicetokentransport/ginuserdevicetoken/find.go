package ginuserdevicetoken

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/userdevicetoken/userdevicetokenbiz"
	"foodlive/modules/userdevicetoken/userdevicetokenstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindUserDeviceToken(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}

		store := userdevicetokenstore.NewSQLStore(appCtx.GetDatabase())
		biz := userdevicetokenbiz.NewFindUserDeviceTokenBiz(store)

		result, err := biz.FindUserDeviceTokenBiz(c.Request.Context(), userIdRaw.(int))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
