package ginuseraddress

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/city/citystore"
	"foodlive/modules/useraddress/useraddressbiz"
	"foodlive/modules/useraddress/useraddressmodel"
	"foodlive/modules/useraddress/useraddressstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateUserAddress(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data useraddressmodel.UserAddressUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		updateUserAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())
		cityStore := citystore.NewSqlStore(appCtx.GetDatabase())
		updateUserAddressBiz := useraddressbiz.NewUpdateUserAddressBiz(cityStore, updateUserAddressStore)

		if err := updateUserAddressBiz.UpdateUserAddressBiz(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
