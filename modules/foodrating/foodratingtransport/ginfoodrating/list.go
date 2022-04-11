package ginfoodrating

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/foodrating/foodratingbiz"
	"foodlive/modules/foodrating/foodratingmodel"
	"foodlive/modules/foodrating/foodratingstore"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListUserRatingFood(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		foodId, err := strconv.Atoi(c.Param("food_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var filter foodratingmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		foodRatingStore := foodratingstore.NewSQLStore(appCtx.GetDatabase())
		biz := foodratingbiz.NewlistUserRatingFood(foodRatingStore)

		result, err := biz.ListRatingFoodBiz(c.Request.Context(), foodId, &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}



func ListMyRatingFood(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var filter foodratingmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userId := c.MustGet(common.KeyUserHeader).(int)

		foodRatingStore := foodratingstore.NewSQLStore(appCtx.GetDatabase())
		biz := foodratingbiz.NewlistUserRatingFood(foodRatingStore)

		result, err := biz.LisMytRatingFoodBiz(c.Request.Context(), userId, &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSuccessResponse(result, paging, filter))
	}
}
