package ginrestaurant

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantbiz"
	"foodlive/modules/restaurant/restaurantmodel"
	"foodlive/modules/restaurant/restaurantrepo"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/restaurantlike/restaurantlikestore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}

		paging.Fulfill()

		store := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		likeStore := restaurantlikestore.NewSQLStore(appCtx.GetDatabase())
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		userId := c.MustGet(common.KeyUserHeader).(int)

		result, err := biz.ListRestaurant(c.Request.Context(), userId, &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			if i == len(result)-1 {
				paging.NextCursor = result[i].Id
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListRestaurantOwner(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userId := c.MustGet(common.KeyUserHeader).(int)

		store := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		likeStore := restaurantlikestore.NewSQLStore(appCtx.GetDatabase())
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurantOwner(c.Request.Context(), userId, &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			if i == len(result)-1 {
				paging.NextCursor = result[i].Id
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListRestaurantForAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}

		paging.Fulfill()

		store := restaurantstore.NewSqlStore(appCtx.GetDatabase())
		likeStore := restaurantlikestore.NewSQLStore(appCtx.GetDatabase())
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurantForAdmin(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			if i == len(result)-1 {
				paging.NextCursor = result[i].Id
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
