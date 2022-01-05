package ginrestaurantlike

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/restaurantlike/restaurantlikebiz"
	"fooddelivery/modules/restaurantlike/restaurantlikestore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserUnLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userId := c.MustGet(common.KeyUserHeader).(int)

		store := restaurantlikestore.NewSQLStore(appCtx.GetDatabase())
		biz := restaurantlikebiz.NewUserUnLikeRestaurantBiz(store, appCtx.GetPubSubProvider())

		err = biz.UnLikeRestaurant(c.Request.Context(), userId, id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(true))
	}
}
