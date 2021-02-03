package repository

import (
	"currency_exchange_collector/database/model"

	"github.com/jinzhu/gorm"
)

type CurrencyExchangeRunsRepository interface {
	Save(entity *model.CurrencyExchangeRun) error
	GetLast() (*model.CurrencyExchangeRun, error)
	BeginTransaction() *gorm.DB
}
