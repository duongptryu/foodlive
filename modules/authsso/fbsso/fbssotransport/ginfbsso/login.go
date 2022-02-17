package ginfbsso

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/authsso/fbsso/fbssobiz"
	"foodlive/modules/authsso/fbsso/fbssomodel"
	"foodlive/modules/authsso/fbsso/fbssostore"
	"foodlive/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func UserFacebookLogin(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data fbssomodel.FacebookJwtInput

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		fbRegisterStore := fbssostore.NewAuthSsoStore()

		fbLoginBiz := fbssobiz.NewLoginFbBiz(fbRegisterStore, userStore, appCtx.GetTokenProvider(), 60*60*24*30)
		account, err := fbLoginBiz.LoginFbBiz(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(account))
	}
}
