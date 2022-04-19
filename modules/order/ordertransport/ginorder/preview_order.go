package ginorder

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/cart/cartstore"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/useraddress/useraddressstore"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PreviewOrder(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		address_id, err := strconv.Atoi(c.Query("address_id"))
		if err != nil {
			panic(common.ErrParseJson(common.NewCustomError(nil, "Cannot parse address_id", "ErrParseAddressId")))
		}

		var data ordermodel.OrderPreviewReq
		data.AddressId = address_id

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userIdInt := userIdRaw.(int)

		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		userAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())

		orderBiz := orderbiz.NewPreviewOrder(cartStore, restaurantStore, userAddressStore)

		resp, err := orderBiz.PreviewOrderBiz(c.Request.Context(), userIdInt, &data, appCtx.GetCryptoPaymentProvider())
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(resp))
	}
}
