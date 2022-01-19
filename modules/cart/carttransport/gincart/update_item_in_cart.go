package gincart

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/cart/cartbiz"
	"foodlive/modules/cart/cartmodel"
	"foodlive/modules/cart/cartstore"
	"github.com/gin-gonic/gin"
)

func UpdateItemInCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data cartmodel.CartItemUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}
		data.UserId = userIdRaw.(int)

		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())
		biz := cartbiz.NewUpdateFoodInCartBiz(cartStore)

		if err := biz.UpdateFoodInCartBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
