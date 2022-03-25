package ginorder

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/orderdetail/orderdetailstore"
	"foodlive/modules/ordertracking/ordertrackingstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FindOrder(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderId, err := strconv.Atoi(c.Param("order_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userId := userIdRaw.(int)

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
		orderDetail := orderdetailstore.NewSqlStore(appCtx.GetDatabase())
		orderTracking := ordertrackingstore.NewSqlStore(appCtx.GetDatabase())

		listOrderBiz := orderbiz.NewFindOrderBiz(orderStore, orderDetail, orderTracking)

		result, err := listOrderBiz.FindOrderBiz(c.Request.Context(), orderId, userId)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}

func FindOrderCryptoInWeb(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderId, err := strconv.Atoi(c.Param("order_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
		orderDetail := orderdetailstore.NewSqlStore(appCtx.GetDatabase())
		orderTracking := ordertrackingstore.NewSqlStore(appCtx.GetDatabase())

		listOrderBiz := orderbiz.NewFindOrderBiz(orderStore, orderDetail, orderTracking)

		result, err := listOrderBiz.FindOrderCryptoBiz(c.Request.Context(), orderId)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}
