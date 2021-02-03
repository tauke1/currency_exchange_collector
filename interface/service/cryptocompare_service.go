package service

type CryptoCompareErrorResponse struct {
	Response string      `json:"Response"`
	Message  string      `json:"Message"`
	Type     int32       `json:"Type"`
	Data     interface{} `json:"Data"`
}

type CryptoComparePriceMultifullResponse struct {
	Raw     map[string]map[string]CryptoComparePriceMultifullResponseRawItem     `json:"RAW"`
	Display map[string]map[string]CryptoComparePriceMultifullResponseDisplayItem `json:"DISPLAY"`
}

type CryptoComparePriceMultifullResponseRawItem struct {
	Change24Hour         float64 `json:"CHANGE24HOUR"`
	ChangePCT24Hour      float64 `json:"CHANGEPCT24HOUR"`
	Open24Hour           float64 `json:"OPEN24HOUR"`
	Volume24Hour         float64 `json:"VOLUME24HOUR"`
	Volume24HourTo       float64 `json:"VOLUME24HOURTO"`
	Low24Hour            float64 `json:"LOW24HOUR"`
	High24Hour           float64 `json:"HIGH24HOUR"`
	Price                float64 `json:"PRICE"`
	Supply               float64 `json:"SUPPLY"`
	MarketCapitalization float64 `json:"MKTCAP"`
}

type CryptoComparePriceMultifullResponseDisplayItem struct {
	Change24Hour         string `json:"CHANGE24HOUR"`
	ChangePCT24Hour      string `json:"CHANGEPCT24HOUR"`
	Open24Hour           string `json:"OPEN24HOUR"`
	Volume24Hour         string `json:"VOLUME24HOUR"`
	Volume24HourTo       string `json:"VOLUME24HOURTO"`
	Low24Hour            string `json:"LOW24HOUR"`
	High24Hour           string `json:"HIGH24HOUR"`
	Price                string `json:"PRICE"`
	Supply               string `json:"SUPPLY"`
	MarketCapitalization string `json:"MKTCAP"`
}

type CryptoCompareService interface {
	GetPriceMultifull(fsyms []string, tsyms []string) (*CryptoComparePriceMultifullResponse, error)
}
