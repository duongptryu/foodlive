package ginrestaurant

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantbiz"
	"foodlive/modules/restaurant/restaurantrepo"
	"foodlive/modules/restaurant/restaurantstore"
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
		//likeStore := restaurantlikestorage.NewSQLStore(appCtx.GetUploadProvider())
		repo := restaurantrepo.NewFindRestaurantRepo(store)
		biz := restaurantbiz.NewFindRestaurantBiz(repo)

		result, err := biz.FindRestaurantById(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
