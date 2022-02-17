package orderstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/order/ordermodel"
)


func (s *sqlStore) ListOrder(ctx context.Context,
	condition map[string]interface{},
	filter *ordermodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]ordermodel.Order, error) {
	var result []ordermodel.Order

	db := s.db

	db = db.Table(ordermodel.Order{}.TableName()).Where(condition)

	// if v := filter; v != nil {

	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if paging.FakeCursor > 0 {
		db = db.Where("id < ?", paging.FakeCursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
