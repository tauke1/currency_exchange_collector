package model

import "time"

type CurrencyExchangeRun struct {
	ID                uint `gorm:"primary_key;auto_increment"`
	Date              time.Time
	CurrencyExchanges []CurrencyExchange
}

func (CurrencyExchangeRun) TableName() string {
	return "currency_exchange_runs"
}
