package ginorder

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/restaurant/restaurantstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ListOrderRestaurant(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		rstId, err := strconv.Atoi(c.Param("restaurant_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}
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
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())

		listOrderBiz := orderbiz.NewListOrderRestaurantBiz(orderStore, restaurantStore)

		result, err := listOrderBiz.ListOrderRestaurantBiz(c.Request.Context(), userId, rstId, &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListCurrentOrderRestaurant(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		rstId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

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
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())

		listOrderBiz := orderbiz.NewListOrderRestaurantBiz(orderStore, restaurantStore)

		result, err := listOrderBiz.ListCurrentOrderRestaurantBiz(c.Request.Context(), userId, rstId, &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}