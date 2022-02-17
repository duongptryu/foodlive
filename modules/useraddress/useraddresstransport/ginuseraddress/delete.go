package ginuseraddress

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/useraddress/useraddressbiz"
	"foodlive/modules/useraddress/useraddressstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteUserAddress(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		deleteUserAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())
		deleteUserAddressBiz := useraddressbiz.NewDeleteUserAddressBiz(deleteUserAddressStore)

		if err := deleteUserAddressBiz.DeleteUserAddressBiz(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
