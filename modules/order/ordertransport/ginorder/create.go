package ginorder

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/cart/cartstore"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"github.com/gin-gonic/gin"
)

func CreateOrder(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data ordermodel.OrderCreate

		if err := c.BindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userIdInt := userIdRaw.(int)

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())

		orderBiz := orderbiz.NewCreateOrderBiz(orderStore, cartStore, appCtx.GetPaymentProvider())

		resp, err := orderBiz.CreateOrderBiz(c.Request.Context(), userIdInt, &data)
		if err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(resp))
	}
}
