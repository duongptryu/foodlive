package statisticbiz

import (
	"context"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/statistic/statisticmodel"
)

type statsFoodBiz struct {
	foodStore foodstore.FoodStore
}

func NewStatsFoodBiz(foodStore foodstore.FoodStore) *statsFoodBiz {
	return &statsFoodBiz{
		foodStore: foodStore,
	}
}

func (biz *statsFoodBiz) StatsFoodBiz(ctx context.Context) (*statisticmodel.StatsFoodByCate, error) {
	foods, err := biz.foodStore.ListFoodWithoutPaging(ctx, map[string]interface{}{"status": true}, nil, "Category")
	if err != nil {
		return nil, err
	}

	var tmpMap = make(map[string]int)
	var cate []string
	var value []int

	for _, v := range foods {
		var key string
		if cate := v.Category; cate == nil {
			continue
		} else {
			key = cate.Name
		}

		if _, exist := tmpMap[key]; exist {
			tmpMap[key] = tmpMap[key] + 1
		} else {
			tmpMap[key] = 1
		}
	}

	for k, v := range tmpMap {
		cate = append(cate, k)
		value = append(value, v)
	}

	return &statisticmodel.StatsFoodByCate{
		value,
		cate,
	}, nil
}
