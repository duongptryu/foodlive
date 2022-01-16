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

func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userId := c.MustGet(common.KeyUserHeader).(int)

		data := restaurantlikemodel.Like{
			RestaurantId: id,
			UserId:       userId,
		}

		store := restaurantlikestore.NewSQLStore(appCtx.GetDatabase())
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubSubProvider())

		err = biz.LikeRestaurant(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(true))
	}
}
