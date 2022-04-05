package ginuser

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/user/userbiz"
	"foodlive/modules/user/userstorage"
	"foodlive/modules/useraddress/useraddressstore"
	"foodlive/modules/userdevicetoken/userdevicetokenstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AdminFindUser(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		userAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())
		userDeviceToken := userdevicetokenstore.NewSQLStore(appCtx.GetDatabase())
		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())

		biz := userbiz.NewAdminFindUserBiz(userStore, userAddressStore, userDeviceToken, orderStore)
		result, err := biz.AdminFindUserBiz(c.Request.Context(), userId)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}

func GetUserProfile(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		userId := c.MustGet(common.KeyUserHeader).(int)
		userStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		userAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())
		userDeviceToken := userdevicetokenstore.NewSQLStore(appCtx.GetDatabase())
		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())

		biz := userbiz.NewAdminFindUserBiz(userStore, userAddressStore, userDeviceToken, orderStore)
		result, err := biz.AdminFindUserBiz(c.Request.Context(), userId)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}
