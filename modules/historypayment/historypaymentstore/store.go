package historypaymentstore

import (
	"context"
	"foodlive/modules/historypayment/historypaymentmodel"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type HistoryPaymentStore interface {
	CreateHistoryPayment(ctx context.Context, data *historypaymentmodel.HistoryPaymentCreate) error
}
