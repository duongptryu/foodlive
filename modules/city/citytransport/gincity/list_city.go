package gincity

import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/city/citybiz"
	"foodlive/modules/city/citystore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCity(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cityStore := citystore.NewSqlStore(appCtx.GetDatabase())
		biz := citybiz.NewListCityBiz(cityStore)

		result, err := biz.ListCityBiz(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
