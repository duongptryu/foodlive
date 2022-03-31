package ginstatistic

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/statistic/statisticbiz"
	"foodlive/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func GetOverview(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		userStore := userstorage.NewSQLStore(appCtx.GetDatabase())
		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
		rstStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())

		biz := statisticbiz.NewOverviewBiz(orderStore, userStore, rstStore, appCtx.GetCryptoPaymentProvider())

		result, err := biz.OverviewBiz(c.Request.Context());
		if  err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}