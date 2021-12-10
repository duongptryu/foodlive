package ginfbsso

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/authsso/fbsso/fbssobiz"
	"fooddelivery/modules/authsso/fbsso/fbssomodel"
	"fooddelivery/modules/authsso/fbsso/fbssostore"
	"fooddelivery/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func UserFacebookRegister(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data fbssomodel.FacebookJwtInput

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		fbRegisterStore := fbssostore.NewAuthSsoStore()

		fbRegisterBiz := fbssobiz.NewRegisterFbSsoBiz(fbRegisterStore, userStore)
		result, err := fbRegisterBiz.RegisterFbSsoBiz(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(result))
	}
}
