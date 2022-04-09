package ginfoodrating

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/foodrating/foodratingbiz"
	"foodlive/modules/foodrating/foodratingmodel"
	"foodlive/modules/foodrating/foodratingstore"

	"github.com/gin-gonic/gin"
)


func UserRatingFood(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data foodratingmodel.FoodRatingCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		data.UserId = userIdRaw.(int)

		foodRatingStore := foodratingstore.NewSQLStore(appCtx.GetDatabase())
		foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())
		biz := foodratingbiz.NewUserRatingFoodBiz(foodStore, foodRatingStore, appCtx.GetPubSubProvider())

		if err := biz.UserRatingFoodBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
