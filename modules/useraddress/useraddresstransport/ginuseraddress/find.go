package ginuseraddress

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/useraddress/useraddressbiz"
	"foodlive/modules/useraddress/useraddressstore"
	"github.com/gin-gonic/gin"
)

func FindDefaultUserAddress(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := c.MustGet(common.KeyUserHeader).(int)

		userAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())
		biz := useraddressbiz.NewFindUserAddressBiz(userAddressStore)

		result, err := biz.FindDefaultUserAddressBiz(c.Request.Context(), userId)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}
