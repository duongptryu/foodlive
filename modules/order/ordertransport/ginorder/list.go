package ginorder

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"github.com/gin-gonic/gin"
)

func ListOrder(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var filter ordermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userId := userIdRaw.(int)

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())

		listOrderBiz := orderbiz.NewListOrderBiz(orderStore)

		result, err := listOrderBiz.ListOrderBiz(c.Request.Context(), userId, &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListMyCurrentOrder(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var filter ordermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userId := userIdRaw.(int)

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())

		listOrderBiz := orderbiz.NewListOrderBiz(orderStore)

		result, err := listOrderBiz.ListMyCurrentOrderBiz(c.Request.Context(), userId, &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}
