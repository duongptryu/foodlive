package ginstatistic

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/statistic/statisticbiz"
	"foodlive/modules/user/userstorage"
	"github.com/gin-gonic/gin"
)

func GetStatsUser(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		userStore := userstorage.NewSQLStore(appCtx.GetDatabase())

		biz := statisticbiz.NewStatsUserBiz(userStore)

		result, err := biz.StatsUserBiz(c.Request.Context(), 2022)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(result))
	}
}
