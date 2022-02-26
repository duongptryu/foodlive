package ginorder

import (
	"encoding/json"
	"fmt"
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"github.com/gin-gonic/gin"
)

func HandleWebHookPayment(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Requset from momo", c.Request)
		c.Request.ParseForm()
		var result = make(map[string]string)
		for key, value := range c.Request.PostForm {
			result[key] = value[0]
		}
		//decode
		var resp ordermodel.WebHookPayment
		respByte, err := json.Marshal(result)
		if err != nil {
			panic(common.ErrParseJson(err))
		}
		err = json.Unmarshal(respByte, &resp)
		if err != nil {
			return
		}

		orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())

		biz := orderbiz.NewWebHookPaymentBiz(orderStore)

		if err := biz.WebHookPaymentBiz(c.Request.Context(), &resp); err != nil {
			panic(err)
		}

		c.JSON(200, nil)
	}
}
