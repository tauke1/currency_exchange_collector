package controller

import (
	"currency_exchange_collector/interface/controller"
)

type AppController struct {
	CurrenciesExchangeController controller.CurrenciesExchangeController
}

func NewAppController(currenciesExchangeController controller.CurrenciesExchangeController) *AppController {
	if currenciesExchangeController == nil {
		panic("currenciesExchangeController argument must not be nil")
	}

	controller := new(AppController)
	controller.CurrenciesExchangeController = currenciesExchangeController
	return controller
}
