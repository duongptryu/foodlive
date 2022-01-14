package citystore

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/modules/city/citymodel"
)

func (s *sqlStore) UpdateCity(ctx context.Context, id int, data *citymodel.CityUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}