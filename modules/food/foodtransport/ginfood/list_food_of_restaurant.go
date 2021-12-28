package ginfood

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/food/foodbiz"
	"fooddelivery/modules/food/foodmodel"
	"fooddelivery/modules/food/foodstore"
	"fooddelivery/modules/restaurant/restaurantstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ListFoodOfRestaurant(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		rId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var filter foodmodel.Filter
		if err := c.ShouldBindJSON(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		listFoodBiz := foodbiz.NewListFoodOfRestaurantBiz(foodStore, restaurantStore)

		result, err := listFoodBiz.ListFoodOfRestaurantBiz(c.Request.Context(), rId, &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}
