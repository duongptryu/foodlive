package ginfood

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/category/categorystore"
	"foodlive/modules/food/foodbiz"
	"foodlive/modules/food/foodmodel"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/restaurant/restaurantstore"
	"github.com/gin-gonic/gin"
)

func CreateFood(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data foodmodel.FoodCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userId := userIdRaw.(int)

		foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		categoryStore := categorystore.NewSqlStore(appCtx.GetDatabase())
		foodBiz := foodbiz.NewCreateFoodBiz(foodStore, restaurantStore, categoryStore)

		if err := foodBiz.CreateFoodBiz(c.Request.Context(), &data, userId); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
