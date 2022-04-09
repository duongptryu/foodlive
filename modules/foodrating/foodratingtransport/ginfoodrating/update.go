package ginfoodrating

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/foodrating/foodratingbiz"
	"foodlive/modules/foodrating/foodratingmodel"
	"foodlive/modules/foodrating/foodratingstore"
	"strconv"

	"github.com/gin-gonic/gin"
)


func UpdateFoodRating(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		rstid, err := strconv.Atoi(c.Param("id_rating"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data foodratingmodel.FoodRatingUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userId := c.MustGet(common.KeyUserHeader).(int)

		foodRatingStore := foodratingstore.NewSQLStore(appCtx.GetDatabase())
		foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())
		biz := foodratingbiz.NewUserUpdateRatingFoodBiz(foodStore, foodRatingStore, appCtx.GetPubSubProvider())

		if err := biz.UpdateRestaurantRatingBiz(c.Request.Context(), rstid, userId, &data); err != nil {
			panic(err)
		}
		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
