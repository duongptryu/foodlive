package historypaymentstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/historypayment/historypaymentmodel"
)

func (s *sqlStore) UpdateHistoryPayment(ctx context.Context, id int, data *historypaymentmodel.HistoryPaymentUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
