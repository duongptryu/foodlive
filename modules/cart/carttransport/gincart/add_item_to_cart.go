package gincart

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/cart/cartbiz"
	"foodlive/modules/cart/cartmodel"
	"foodlive/modules/cart/cartstore"
	"github.com/gin-gonic/gin"
)

func CreateItemInCart(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data cartmodel.CartItemCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}
		data.UserId = userIdRaw.(int)

		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())
		biz := cartbiz.NewAddFoodToCartBiz(cartStore)

		if err := biz.AddFoodToCartBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
