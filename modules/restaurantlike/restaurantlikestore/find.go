package restaurantlikestore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/restaurantlike/restaurantlikemodel"
)

func (s sqlStore) FindUserLikeRst(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantlikemodel.Like, error) {
	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(condition)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result restaurantlikemodel.Like

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
