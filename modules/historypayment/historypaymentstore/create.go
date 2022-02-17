package historypaymentstore

import (
	"context"
	"foodlive/common"
	"foodlive/modules/historypayment/historypaymentmodel"
)

func (s *sqlStore) CreateHistoryPayment(ctx context.Context, data *historypaymentmodel.HistoryPaymentCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
