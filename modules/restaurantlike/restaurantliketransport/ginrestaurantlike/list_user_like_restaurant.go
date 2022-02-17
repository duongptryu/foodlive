package ginrestaurantlike

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurantlike/restaurantlikebiz"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
	"foodlive/modules/restaurantlike/restaurantlikestore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListUserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: id,
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}

		paging.Fulfill()

		store := restaurantlikestore.NewSQLStore(appCtx.GetDatabase())
		biz := restaurantlikebiz.NewListUserLikeRestaurant(store)

		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
