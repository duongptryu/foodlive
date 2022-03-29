package ginorder

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/ordertracking/ordertrackingstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RstConfirmPrepareDone(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderId, err := strconv.Atoi(c.Param("order_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		ownerId := c.MustGet(common.KeyUserHeader)
		if err := ownerId.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userIdInt := ownerId.(int)

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
		orderTracking := ordertrackingstore.NewSqlStore(appCtx.GetDatabase())

		biz := orderbiz.NewUpdateOrderBiz(orderStore, orderTracking)

		err = biz.RstConfirmPrepareDone(c.Request.Context(), orderId, userIdInt)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}

func UserConfirmReceived(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderId, err := strconv.Atoi(c.Param("order_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userId := c.MustGet(common.KeyUserHeader)
		if err := userId.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userIdInt := userId.(int)

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
		orderTracking := ordertrackingstore.NewSqlStore(appCtx.GetDatabase())

		biz := orderbiz.NewUpdateOrderBiz(orderStore, orderTracking)

		err = biz.UserConfirmReceived(c.Request.Context(), orderId, userIdInt)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}