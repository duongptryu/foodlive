package gingooglesso

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/authsso/googlesso/googlessobiz"
	"foodlive/modules/authsso/googlesso/googlessomodel"
	"foodlive/modules/authsso/googlesso/googlessostore"
	"foodlive/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func UserGoogleRegister(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data googlessomodel.GoogleJwtInput

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		googleRegisterStore := googlessostore.NewAuthSsoStore()

		ggRegisterBiz := googlessobiz.NewRegisterGoogleBiz(googleRegisterStore, userStore)
		result, err := ggRegisterBiz.RegisterGoogleBiz(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(result))
	}
}
