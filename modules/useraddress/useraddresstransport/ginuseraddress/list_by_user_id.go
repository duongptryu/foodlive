package ginuseraddress

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/useraddress/useraddressbiz"
	"foodlive/modules/useraddress/useraddressstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUserAddress(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdRaw := c.MustGet(common.KeyUserHeader)

		listUserAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())
		listUserAddressBiz := useraddressbiz.NewListUserAddressBiz(listUserAddressStore)

		result, err := listUserAddressBiz.ListUserAddressBiz(c.Request.Context(), userIdRaw.(int))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
