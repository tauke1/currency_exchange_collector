package main

import (
	"currency_exchange_collector/registry"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	fmt.Println("Running server!")
	registry.RegisterRoutes(r)
	r.Run()
}
