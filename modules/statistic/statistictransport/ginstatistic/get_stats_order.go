package ginstatistic

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/statistic/statisticbiz"
	"github.com/gin-gonic/gin"
)

func GetStatsOrder(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())

		biz := statisticbiz.NewStatsOrderBiz(orderStore)

		result, err := biz.StatsOrderBiz(c.Request.Context(), 2022)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}
