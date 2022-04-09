package ginfoodlike

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/foodlike/foodlikebiz"
	"foodlive/modules/foodlike/foodlikemodel"
	"foodlive/modules/foodlike/foodlikestore"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListUserLikeFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("food_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		filter := foodlikemodel.Filter{
			FoodId: id,
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}

		paging.Fulfill()

		store := foodlikestore.NewSQLStore(appCtx.GetDatabase())
		biz := foodlikebiz.NewlistUserLikeFood(store)

		result, err := biz.ListUserLikeFood(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListMyLikeFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter foodlikemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userId := c.MustGet(common.KeyUserHeader).(int)

		store := foodlikestore.NewSQLStore(appCtx.GetDatabase())
		biz := foodlikebiz.NewlistUserLikeFood(store)

		result, err := biz.ListMyLikeFood(c.Request.Context(), userId, &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
