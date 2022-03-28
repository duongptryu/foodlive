package ginorder

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/cart/cartstore"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/orderdetail/orderdetailstore"
	"foodlive/modules/ordertracking/ordertrackingstore"
	"foodlive/modules/useraddress/useraddressstore"
	"github.com/gin-gonic/gin"
)

func CreateOrderMomo(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var checkOut ordermodel.Checkout
		if err := c.BindJSON(&checkOut); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userIdInt := userIdRaw.(int)

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())
		orderDetail := orderdetailstore.NewSqlStore(appCtx.GetDatabase())
		orderTracking := ordertrackingstore.NewSqlStore(appCtx.GetDatabase())
		userAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())

		orderBiz := orderbiz.NewCreateOrderBiz(orderStore, orderDetail, orderTracking, userAddressStore, cartStore, appCtx.GetPaymentProvider())

		resp, err := orderBiz.CreateOrderMomoBiz(c.Request.Context(), userIdInt, &checkOut)
		if err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(resp))
	}
}

func CreateOrderCrypto(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var checkOut ordermodel.Checkout
		if err := c.BindJSON(&checkOut); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userIdInt := userIdRaw.(int)

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
		cartStore := cartstore.NewSqlStore(appCtx.GetDatabase())
		orderDetail := orderdetailstore.NewSqlStore(appCtx.GetDatabase())
		orderTracking := ordertrackingstore.NewSqlStore(appCtx.GetDatabase())
		userAddressStore := useraddressstore.NewSQLStore(appCtx.GetDatabase())

		orderBiz := orderbiz.NewCreateOrderBiz(orderStore, orderDetail, orderTracking, userAddressStore, cartStore, appCtx.GetPaymentProvider())

		resp, err := orderBiz.CreateOrderCryptoBiz(c.Request.Context(), userIdInt, &checkOut)
		if err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(resp))
	}
}
