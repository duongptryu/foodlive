package ginfoodlike

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/foodlike/foodlikebiz"
	"foodlive/modules/foodlike/foodlikemodel"
	"foodlive/modules/foodlike/foodlikestore"
	"net/http"
	"github.com/gin-gonic/gin"
)


func UserLikeFood(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data foodlikemodel.FoodLikeCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrParseJson(err))
		}
		userId := c.MustGet(common.KeyUserHeader).(int)

		data.UserId = userId

		store := foodlikestore.NewSQLStore(appCtx.GetDatabase())
		biz := foodlikebiz.NewUserLikeFoodBiz(store, appCtx.GetPubSubProvider())

		err := biz.UserLikeFoodBiz(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(true))
	}
}
