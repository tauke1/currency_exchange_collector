package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type CurrenciesExchangeController interface {
	GetPriceMultifull(context *gin.Context)
	RefreshPriceMultifull(context *gin.Context)
	OnConnectingToWSChannel(s *melody.Session)
}
