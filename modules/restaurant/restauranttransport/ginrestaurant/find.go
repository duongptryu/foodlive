package ginrestaurant

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantbiz"
	"foodlive/modules/restaurant/restaurantrepo"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/restaurantlike/restaurantlikestore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		store := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		likeStore := restaurantlikestore.NewSQLStore(appCtx.GetDatabase())
		repo := restaurantrepo.NewFindRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewFindRestaurantBiz(repo)

		userId := c.MustGet(common.KeyUserHeader).(int)

		result, err := biz.FindRestaurantById(c.Request.Context(), id, userId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}

func FindRestaurantWithoutStatus(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		store := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		likeStore := restaurantlikestore.NewSQLStore(appCtx.GetDatabase())
		repo := restaurantrepo.NewFindRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewFindRestaurantBiz(repo)

		result, err := biz.FindRestaurantByIdWithoutStatus(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
