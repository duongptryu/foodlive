package ginuser

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/user/userbiz"
	"foodlive/modules/user/usermodel"
	"foodlive/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AdminUpdateUser(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data usermodel.UserUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		biz := userbiz.NewUpdateUserBiz(userStore)

		if err := biz.AdminUpdateUserBiz(c.Request.Context(), userId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}

func AccountSSOUpdate(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userId := c.MustGet(common.KeyUserHeader).(int)

		userStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		biz := userbiz.NewUpdateUserBiz(userStore)

		if err := biz.UpdateUserBiz(c.Request.Context(), userId, &data, appCtx.GetMyCache(), appCtx.GetMySms()); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
