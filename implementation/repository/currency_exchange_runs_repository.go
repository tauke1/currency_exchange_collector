package repository

import (
	"currency_exchange_collector/database/model"

	"github.com/jinzhu/gorm"
)

type currencyExchangeRunsRepository struct {
	db *gorm.DB
}

func (repo *currencyExchangeRunsRepository) Save(entity *model.CurrencyExchangeRun) error {
	if entity == nil {
		panic("entity argument must not be nil")
	}

	return repo.db.Save(entity).Error
}

func (repo *currencyExchangeRunsRepository) GetLast() (*model.CurrencyExchangeRun, error) {
	entity := model.CurrencyExchangeRun{}
	err := repo.db.Last(&entity).Error
	return &entity, err
}

func (repo *currencyExchangeRunsRepository) BeginTransaction() *gorm.DB {
	return repo.db.Begin()
}

func NewCcurrencyExchangeRunsRepository(db *gorm.DB) *currencyExchangeRunsRepository {
	if db == nil {
		panic("db argument must not be nil")
	}

	repo := new(currencyExchangeRunsRepository)
	repo.db = db
	return repo
}
