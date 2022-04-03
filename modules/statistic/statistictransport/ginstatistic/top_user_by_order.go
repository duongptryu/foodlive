package ginstatistic

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/statistic/statisticbiz"
	"github.com/gin-gonic/gin"
)

func GetStatsTopUserByOrder(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())

		biz := statisticbiz.NewStatsUserByOrderBiz(orderStore)

		result, err := biz.StatsUserByOrderBiz(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}
