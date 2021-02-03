package repository

import (
	"currency_exchange_collector/database/model"

	"github.com/jinzhu/gorm"
)

type currencyExchangesRepository struct {
	db *gorm.DB
}

func (repo *currencyExchangesRepository) Save(entity *model.CurrencyExchange) error {
	if entity == nil {
		panic("entity argument must not be nil")
	}

	return repo.db.Save(entity).Error
}

func (repo *currencyExchangesRepository) GetListByRunID(runID int) ([]model.CurrencyExchange, error) {
	entities := make([]model.CurrencyExchange, 0)
	err := repo.db.Where("currency_exchange_run_id = ?", runID).Find(&entities).Error
	return entities, err
}

func (repo *currencyExchangesRepository) BeginTransaction() *gorm.DB {
	return repo.db.Begin()
}

func NewCurrencyExchangesRepository(db *gorm.DB) *currencyExchangesRepository {
	if db == nil {
		panic("db argument must not be nil")
	}

	repo := new(currencyExchangesRepository)
	repo.db = db
	return repo
}
