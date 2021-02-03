package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type configuration struct {
	DbHost               string
	DbPassword           string
	DbUser               string
	DbName               string
	ServerPort           int
	CryptoCompareBaseUrl string
	FSyms                []string
	TSyms                []string
	FSymsMap             map[string]bool
	TSymsMap             map[string]bool
}

var C configuration

func ReadConfig() {
	Config := &C
	viper.SetConfigType("yml")
	viper.SetConfigFile("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Config.FSymsMap = make(map[string]bool)
	for _, currency := range Config.FSyms {
		Config.FSymsMap[currency] = true
	}

	Config.TSymsMap = make(map[string]bool)
	for _, currency := range Config.TSyms {
		Config.TSymsMap[currency] = true
	}
}
