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
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data foodmodel.FoodUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err != 0 && userIdRaw == nil {
			panic(common.ErrUnAuthorization)
		}
		userId := userIdRaw.(int)

		foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())
		restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		categoryStore := categorystore.NewSqlStore(appCtx.GetDatabase())
		updateFoodBiz := foodbiz.NewUpdateFoodBiz(foodStore, restaurantStore, categoryStore)

		if err := updateFoodBiz.UpdateFoodBiz(c.Request.Context(), id, userId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
