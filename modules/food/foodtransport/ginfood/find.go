package ginfood

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/food/foodbiz"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/foodlike/foodlikestore"
	"foodlive/modules/restaurant/restaurantstore"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FindFood(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		foodLikeStore := foodlikestore.NewSQLStore(appCtx.GetDatabase())
		listFoodBiz := foodbiz.NewFindFoodBiz(foodStore, restaurantStore, foodLikeStore)

		userId := c.MustGet(common.KeyUserHeader).(int)

		result, err := listFoodBiz.FindFoodBiz(c.Request.Context(), fId, userId)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}
