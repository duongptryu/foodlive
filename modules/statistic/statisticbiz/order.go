package statisticbiz

import (
	"context"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/statistic/statisticmodel"
	"time"
)

type statsOrderBiz struct {
	orderStore orderstore.OrderStore
}

func NewStatsOrderBiz(orderStore orderstore.OrderStore) *statsOrderBiz {
	return &statsOrderBiz{
		orderStore: orderStore,
	}
}

func (biz *statsOrderBiz) StatsOrderBiz(ctx context.Context, year int) (*statisticmodel.StatsOrder, error) {
	now := time.Now()
	yearCondition := time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())

	filter := ordermodel.Filter{
		CreatedAtGt: &yearCondition,
	}
	result, err := biz.orderStore.ListOrderWithoutPaging(ctx, nil, &filter)
	if err != nil {
		return nil, err
	}

	var momoOrder []int
	var cryptoOrder []int
	var cate []string
	type temp struct {
		momoOrder   int
		cryptoOrder int
	}
	var check = make(map[string]temp)

	for _, v := range result {
		//var key = v.CreatedAt.Format("2006-01-02")
		var key = v.CreatedAt.Month().String()
		if _, exist := check[key]; exist {
			if v.TypePayment == ordermodel.TypeCrypto {
				check[key] = temp{
					momoOrder:   check[key].momoOrder + 1,
					cryptoOrder: check[key].cryptoOrder,
				}
			} else {
				check[key] = temp{
					momoOrder:   check[key].momoOrder,
					cryptoOrder: check[key].cryptoOrder + 1,
				}
			}
		} else {
			if v.TypePayment == ordermodel.TypeCrypto {
				check[key] = temp{
					momoOrder:   1,
					cryptoOrder: 0,
				}
			} else {
				check[key] = temp{
					momoOrder:   0,
					cryptoOrder: 1,
				}
			}
		}
	}

	for d := yearCondition; d.After(time.Now()) == false; d = d.AddDate(0, 1, 0) {
		//k := d.Format("2006-01-02")
		k := d.Month().String()
		if v, exist := check[k]; exist {
			momoOrder = append(momoOrder, v.momoOrder)
			cryptoOrder = append(cryptoOrder, v.cryptoOrder)
		} else {
			momoOrder = append(momoOrder, 0)
			cryptoOrder = append(cryptoOrder, 0)
		}
		cate = append(cate, k)
	}

	return &statisticmodel.StatsOrder{
		momoOrder,
		cryptoOrder,
		cate,
	}, nil
}
