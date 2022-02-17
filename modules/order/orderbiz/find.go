package orderbiz
import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
)

type findOrderBiz struct {
	orderStore orderstore.OrderStore
}

func NewFindOrderBiz(orderStore orderstore.OrderStore) *findOrderBiz {
	return &findOrderBiz{
		orderStore: orderStore,
	}
}

func (biz *findOrderBiz) FindOrderBiz(ctx context.Context, orderId int, userId int) (*ordermodel.Order, error) {
	result, err := biz.orderStore.FindOrder(ctx, map[string]interface{}{"id": orderId, "user_id": userId})
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return result, nil
}