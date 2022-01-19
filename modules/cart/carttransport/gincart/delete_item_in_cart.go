package gincart

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/cart/cartbiz"
	"foodlive/modules/cart/cartstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteAItemInCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId, err := strconv.Atoi(c.Param("food_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}

		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())
		biz := cartbiz.NewDeleteFoodInCartBiz(cartStore)

		if err := biz.DeleteAFoodInCartBiz(c.Request.Context(), userIdRaw.(int), foodId); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}

func DeleteAllItemInCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}

		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())
		biz := cartbiz.NewDeleteFoodInCartBiz(cartStore)

		if err := biz.DeleteAllFoodInCartBiz(c.Request.Context(), userIdRaw.(int)); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
