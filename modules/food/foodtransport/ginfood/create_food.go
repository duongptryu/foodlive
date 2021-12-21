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

func CreateFood(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data foodmodel.FoodCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}
		userId := userIdRaw.(int)
		data.RestaurantId = id

		foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		createFoodBiz := foodbiz.NewCreateFoodBiz(foodStore, restaurantStore)

		if err := createFoodBiz.CreateFoodBiz(c.Request.Context(), &data, userId); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
