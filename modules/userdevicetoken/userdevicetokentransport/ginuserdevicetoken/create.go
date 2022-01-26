package ginuserdevicetoken

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/userdevicetoken/userdevicetokenbiz"
	"foodlive/modules/userdevicetoken/userdevicetokenmodel"
	"foodlive/modules/userdevicetoken/userdevicetokenstore"
	"github.com/gin-gonic/gin"
)

func CreateUserDeviceToken(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data userdevicetokenmodel.UserDeviceTokenCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}
		data.UserId = userIdRaw.(int)

		store := userdevicetokenstore.NewSQLStore(appCtx.GetDatabase())
		biz := userdevicetokenbiz.NewCreateUserDeviceTokenBiz(store)

		if err := biz.CreateUserDeviceTokenBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
