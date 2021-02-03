package service

type GetPriceMultifullCallback func(*CryptoComparePriceMultifullResponse)

type CurrencyExchangesService interface {
	GetPriceMultifull(fsyms []string, tsyms []string, callback GetPriceMultifullCallback) (*CryptoComparePriceMultifullResponse, error)
	RefreshPriceMultifull(callback GetPriceMultifullCallback) (*CryptoComparePriceMultifullResponse, error)
}
