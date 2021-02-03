package repository

import (
	"currency_exchange_collector/database/model"

	"github.com/jinzhu/gorm"
)

type CurrencyExchangesRepository interface {
	Save(entity *model.CurrencyExchange) error
	GetListByRunID(runID int) ([]model.CurrencyExchange, error)
	BeginTransaction() *gorm.DB
}
