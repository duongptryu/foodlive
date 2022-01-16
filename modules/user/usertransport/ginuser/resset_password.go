package ginuser

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/user/userbiz"
	"foodlive/modules/user/usermodel"
	"foodlive/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func UserResetPassword(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserResetPassword

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userActiveStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		userActiveBiz := userbiz.NewUserResetPasswordBiz(userActiveStore, appCtx.GetMyCache())

		if err := userActiveBiz.UserResetPasswordBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
