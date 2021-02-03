package service

import (
	"currency_exchange_collector/config"
	dbModels "currency_exchange_collector/database/model"
	"currency_exchange_collector/interface/repository"
	interfaceService "currency_exchange_collector/interface/service"
	"errors"
	"fmt"
	"time"
)

type currencyExchangeService struct {
	currencyExchangesRepository    repository.CurrencyExchangesRepository
	currencyExchangeRunsRepository repository.CurrencyExchangeRunsRepository
	cryptocompareService           interfaceService.CryptoCompareService
}

func (service *currencyExchangeService) GetPriceMultifull(fsyms []string, tsyms []string, callback interfaceService.GetPriceMultifullCallback) (*interfaceService.CryptoComparePriceMultifullResponse, error) {
	err := service.validateInput(fsyms, tsyms)
	if err != nil {
		return nil, err
	}

	fsymsMap := make(map[string]bool)
	for _, currency := range fsyms {
		fsymsMap[currency] = true
	}

	tsymsMap := make(map[string]bool)
	for _, currency := range tsyms {
		tsymsMap[currency] = true
	}

	// lets request by all fsyms and tsyms from the config, because we need to update whole data
	serviceResp, err := service.cryptocompareService.GetPriceMultifull(config.C.FSyms, config.C.TSyms)
	resp := interfaceService.CryptoComparePriceMultifullResponse{
		Raw:     make(map[string]map[string]interfaceService.CryptoComparePriceMultifullResponseRawItem),
		Display: make(map[string]map[string]interfaceService.CryptoComparePriceMultifullResponseDisplayItem),
	}

	//lets filter serviceResp and fill response with necessary pairs of fsyms and tsyms
	for currency, exchangeRatesMap := range serviceResp.Raw {
		if _, ok := fsymsMap[currency]; ok {
			currencyExchangeMap := make(map[string]interfaceService.CryptoComparePriceMultifullResponseRawItem)
			for exchangeCurrency, value := range exchangeRatesMap {
				if _, ok := tsymsMap[exchangeCurrency]; ok {
					currencyExchangeMap[exchangeCurrency] = value
				}
			}

			resp.Raw[currency] = currencyExchangeMap
		}
	}

	for currency, exchangeRatesMap := range serviceResp.Display {
		if _, ok := fsymsMap[currency]; ok {
			currencyExchangeMap := make(map[string]interfaceService.CryptoComparePriceMultifullResponseDisplayItem)
			for exchangeCurrency, value := range exchangeRatesMap {
				if _, ok := tsymsMap[exchangeCurrency]; ok {
					currencyExchangeMap[exchangeCurrency] = value
				}
			}

			resp.Display[currency] = currencyExchangeMap
		}
	}

	if err != nil {
		return nil, err
	}

	tx := service.currencyExchangesRepository.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	exchangeRun := dbModels.CurrencyExchangeRun{Date: time.Now()}
	err = service.currencyExchangeRunsRepository.Save(&exchangeRun)
	if err != nil {
		return nil, err
	}

	for currency, exchangeRatesMap := range serviceResp.Raw {
		for exchangeCurrency, value := range exchangeRatesMap {
			dbEntity := dbModels.CurrencyExchange{
				FromCurrency:          currency,
				ToCurrency:            exchangeCurrency,
				Change24Hour:          value.Change24Hour,
				ChangePCT24Hour:       value.ChangePCT24Hour,
				Open24Hour:            value.Open24Hour,
				Volume24Hour:          value.Volume24Hour,
				Volume24HourTo:        value.Volume24HourTo,
				Low24Hour:             value.Low24Hour,
				High24Hour:            value.High24Hour,
				Price:                 value.Price,
				Supply:                value.Supply,
				MarketCapitalization:  value.MarketCapitalization,
				CurrencyExchangeRunID: exchangeRun.ID,
			}

			err := service.currencyExchangesRepository.Save(&dbEntity)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	if callback != nil {
		callback(serviceResp)
	}

	return &resp, err
}

func (service *currencyExchangeService) RefreshPriceMultifull(callback interfaceService.GetPriceMultifullCallback) (*interfaceService.CryptoComparePriceMultifullResponse, error) {
	return service.GetPriceMultifull(config.C.FSyms, config.C.TSyms, callback)
}

func (service *currencyExchangeService) validateInput(fsyms []string, tsyms []string) error {
	if fsyms == nil {
		return errors.New("fsyms argument must not be nil")
	} else if len(fsyms) == 0 {
		return errors.New("fsyms argument must not be empty")
	}

	if tsyms == nil {
		return errors.New("tsyms argument must not be nil")
	} else if len(tsyms) == 0 {
		return errors.New("tsyms argument must not be empty")
	}

	for _, currency := range fsyms {
		if _, ok := config.C.FSymsMap[currency]; !ok {
			return errors.New(fmt.Sprint(currency, " is not in allowed fsyms list"))
		}
	}

	for _, currency := range tsyms {
		if _, ok := config.C.TSymsMap[currency]; !ok {
			return errors.New(fmt.Sprint(currency, " is not in allowed tsyms list"))
		}
	}

	return nil
}

func NewCurrencyExchangeService(currencyExchangesRepository repository.CurrencyExchangesRepository,
	cryptocompareService interfaceService.CryptoCompareService, currencyExchangeRunsRepository repository.CurrencyExchangeRunsRepository) *currencyExchangeService {
	if currencyExchangesRepository == nil {
		panic("currencyExchangesRepository argument must not be nil")
	}

	if currencyExchangeRunsRepository == nil {
		panic("currencyExchangeRunsRepository argument must not be nil")
	}

	if cryptocompareService == nil {
		panic("cryptocompareService argument must not be nil")
	}

	service := new(currencyExchangeService)
	service.currencyExchangesRepository = currencyExchangesRepository
	service.currencyExchangeRunsRepository = currencyExchangeRunsRepository
	service.cryptocompareService = cryptocompareService
	return service
}
