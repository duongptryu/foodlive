package gincart

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/cart/cartbiz"
	"foodlive/modules/cart/cartmodel"
	"foodlive/modules/cart/cartstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListItemInCart(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter cartmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}

		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())
		biz := cartbiz.NewListFoodInCartBiz(cartStore)

		result, err := biz.ListFoodInCartBiz(c.Request.Context(), userIdRaw.(int), &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, filter))
	}
}
