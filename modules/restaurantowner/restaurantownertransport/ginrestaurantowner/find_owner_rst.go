package ginrestaurantowner

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurantowner/restaurantownerbiz"
	"foodlive/modules/restaurantowner/restaurantownerstore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindOwnerRstByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		store := restaurantownerstore.NewSqlStore(appCtx.GetDatabase())
		biz := restaurantownerbiz.NewFindOwnerRst(store)

		result, err := biz.FindOwnerRst(c.Request.Context(), ownerId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
