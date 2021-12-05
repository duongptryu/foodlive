package ginuser

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/user/userbiz"
	"fooddelivery/modules/user/usermodel"
	"fooddelivery/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func UserReigster(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userRegisterStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		userRegisterBiz := userbiz.NewRegisterUserBiz(userRegisterStore, appCtx.GetMyCache(), appCtx.GetMySms())

		if err := userRegisterBiz.RegisterUserBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
