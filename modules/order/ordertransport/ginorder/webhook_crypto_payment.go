package ginorder

import (
	"fmt"
	"foodlive/component"
	"foodlive/modules/historypayment/historypaymentstore"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/ordertracking/ordertrackingstore"
	"github.com/gin-gonic/gin"
)

func HandleWebHookCryptoPayment(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		var resp ordermodel.WebHookPayment
		if err := c.ShouldBindJSON(&resp); err != nil {
			fmt.Println(err)
		}

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
		orderTracking := ordertrackingstore.NewSqlStore(appCtx.GetDatabase())
		historyPaymentStore := historypaymentstore.NewSqlStore(appCtx.GetDatabase())
		biz := orderbiz.NewWebHookPaymentBiz(orderStore, orderTracking, historyPaymentStore)

		if err := biz.WebHookPaymentBiz(c.Request.Context(), &resp); err != nil {
			panic(err)
		}

		c.JSON(200, nil)
	}
}
