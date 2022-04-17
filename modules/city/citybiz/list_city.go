package citybiz

import (
	"context"
	"foodlive/modules/city/citymodel"
	"foodlive/modules/city/citystore"
)

type listCityBiz struct {
	cityStore citystore.CityStore
}

func NewListCityBiz(cityStore citystore.CityStore) *listCityBiz {
	return &listCityBiz{
		cityStore: cityStore,
	}
}

func (biz *listCityBiz) ListCityBiz(ctx context.Context) ([]citymodel.City, error) {
	result, err := biz.cityStore.ListCity(ctx, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}
