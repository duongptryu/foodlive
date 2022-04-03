package ginstatistic

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/statistic/statisticbiz"
	"github.com/gin-gonic/gin"
)

func GetStatsFood(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())

		biz := statisticbiz.NewStatsFoodBiz(foodStore)

		result, err := biz.StatsFoodBiz(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}
