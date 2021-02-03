package model

type CurrencyExchange struct {
	ID                    uint   `gorm:"primary_key;auto_increment"`
	FromCurrency          string `gorm:"index:idx_currency_exchanges_from_currency_to_currency"`
	ToCurrency            string `gorm:"index:idx_currency_exchanges_from_currency_to_currency"`
	Change24Hour          float64
	ChangePCT24Hour       float64
	Open24Hour            float64
	Volume24Hour          float64
	Volume24HourTo        float64
	Low24Hour             float64
	High24Hour            float64
	Price                 float64
	Supply                float64
	MarketCapitalization  float64
	CurrencyExchangeRunID uint
	CurrencyExchangeRun   CurrencyExchangeRun `gorm:"foreignKey:CurrencyExchangeRunID"`
}

func (CurrencyExchange) TableName() string {
	return "currency_exchanges"
}
