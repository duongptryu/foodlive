package ginorder
import (
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"github.com/gin-gonic/gin"
)

func CreateOrder(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data ordermodel.OrderCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		data.UserId = userIdRaw.(int)

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())

		orderBiz := orderbiz.NewCreateOrderBiz(orderStore)

		if err := orderBiz.CreateOrderBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(data))
	}
}
