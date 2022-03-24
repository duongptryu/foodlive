package ginorder

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/cart/cartstore"
	"foodlive/modules/order/orderbiz"
	"github.com/gin-gonic/gin"
)

func PreviewOrder(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userIdInt := userIdRaw.(int)

		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())

		orderBiz := orderbiz.NewPreviewOrder(cartStore)

		resp, err := orderBiz.PreviewOrderBiz(c.Request.Context(), userIdInt, appCtx.GetCryptoPaymentProvider())
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(resp))
	}
}
