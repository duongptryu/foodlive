package ginuseraddress

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/city/citystore"
	"foodlive/modules/useraddress/useraddressbiz"
	"foodlive/modules/useraddress/useraddressmodel"
	"foodlive/modules/useraddress/useraddressstore"
	"github.com/gin-gonic/gin"
)

func CreateUserAddress(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data useraddressmodel.UserAddressCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		data.UserId = userIdRaw.(int)

		createUserAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())
		cityStore := citystore.NewSqlStore(appCtx.GetDatabase())
		createUserAddressBiz := useraddressbiz.NewCreateUserAddressBiz(cityStore, createUserAddressStore)

		if err := createUserAddressBiz.CreateUserAddressBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
