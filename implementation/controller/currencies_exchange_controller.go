package controller

import (
	"currency_exchange_collector/interface/service"
	"currency_exchange_collector/model"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type currenciesExchangeController struct {
	CurrencyExchangesService service.CurrencyExchangesService
	Hub                      *melody.Melody
}

func (controller *currenciesExchangeController) GetPriceMultifull(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	tsymsArray := make([]string, 0)
	fsymsArray := make([]string, 0)
	if tsymsQueryArray, ok := queryParams["tsyms"]; ok {
		for _, tsyms := range tsymsQueryArray {
			tsymsArray = append(tsymsArray, strings.Split(tsyms, ",")...)
		}
	} else {
		context.JSON(http.StatusBadRequest, &model.ErrorResponse{Error: "tsyms query parameter is required"})
		return
	}

	if fsymsQueryArray, ok := queryParams["fsyms"]; ok {
		for _, fsyms := range fsymsQueryArray {
			fsymsArray = append(fsymsArray, strings.Split(fsyms, ",")...)
		}
	} else {
		context.JSON(http.StatusBadRequest, &model.ErrorResponse{Error: "fsyms query parameter is required"})
		return
	}

	resp, err := controller.CurrencyExchangesService.GetPriceMultifull(fsymsArray, tsymsArray, controller.WSCallback)
	if err != nil {
		context.JSON(http.StatusBadRequest, &model.ErrorResponse{Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, resp)
}

func (controller *currenciesExchangeController) RefreshPriceMultifull(context *gin.Context) {
	resp, err := controller.CurrencyExchangesService.RefreshPriceMultifull(controller.WSCallback)
	if err != nil {
		context.JSON(http.StatusBadRequest, &model.ErrorResponse{Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, resp)
	return
}

func (controller *currenciesExchangeController) OnConnectingToWSChannel(s *melody.Session) {
	controller.CurrencyExchangesService.RefreshPriceMultifull(controller.WSCallback)
}

func (controller *currenciesExchangeController) WSCallback(resp *service.CryptoComparePriceMultifullResponse) {
	var webResponse interface{}
	if resp == nil {
		webResponse = &model.ErrorResponse{Error: "empty exchange rates found"}
	} else {
		webResponse = resp
	}

	json, _ := json.Marshal(webResponse)
	controller.Hub.Broadcast(json)
}

func NewCurrenciesExchangeController(currencyExchangesService service.CurrencyExchangesService, hub *melody.Melody) *currenciesExchangeController {
	if currencyExchangesService == nil {
		panic("currencyExchangesService argument must not be nil")
	}

	if hub == nil {
		panic("hub argument must not be nil")
	}

	controller := new(currenciesExchangeController)
	controller.CurrencyExchangesService = currencyExchangesService
	controller.Hub = hub
	return controller
}
