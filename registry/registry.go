package registry

import (
	"currency_exchange_collector/config"
	"currency_exchange_collector/database"
	dbModel "currency_exchange_collector/database/model"
	implController "currency_exchange_collector/implementation/controller"
	implRepository "currency_exchange_collector/implementation/repository"
	implService "currency_exchange_collector/implementation/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func registerAppControllerAndMelody() (*implController.AppController, *melody.Melody) {
	config.ReadConfig()
	db := database.NewDB()
	db.AutoMigrate(&dbModel.CurrencyExchangeRun{}, &dbModel.CurrencyExchange{})
	m := melody.New()
	currencyExchangesRepository := implRepository.NewCurrencyExchangesRepository(db)
	currencyExchangeRunsRepository := implRepository.NewCcurrencyExchangeRunsRepository(db)
	cryptoCompareService := implService.NewCryptocompareService(config.C.CryptoCompareBaseUrl)
	currencyExchangeService := implService.NewCurrencyExchangeService(currencyExchangesRepository, cryptoCompareService, currencyExchangeRunsRepository)
	currencyExchangeController := implController.NewCurrenciesExchangeController(currencyExchangeService, m)
	appController := implController.NewAppController(currencyExchangeController)
	m.HandleConnect(appController.CurrenciesExchangeController.OnConnectingToWSChannel)
	return appController, m
}

func RegisterRoutes(engine *gin.Engine) {
	appController, m := registerAppControllerAndMelody()

	engine.GET("price", appController.CurrenciesExchangeController.GetPriceMultifull)
	engine.GET("refresh", appController.CurrenciesExchangeController.RefreshPriceMultifull)
	engine.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
}
