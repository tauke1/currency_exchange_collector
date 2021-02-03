package service

import (
	"bytes"
	interfaceService "currency_exchange_collector/interface/service"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const PriceMultifullAction = "data/pricemultifull"

type cryptoCompareService struct {
	baseAddress string
}

func (service *cryptoCompareService) GetPriceMultifull(fsyms []string, tsyms []string) (*interfaceService.CryptoComparePriceMultifullResponse, error) {
	if fsyms == nil {
		return nil, errors.New("fsyms argument must not be nil")
	} else if len(fsyms) == 0 {
		return nil, errors.New("fsyms argument must not be empty")
	}

	if tsyms == nil {
		return nil, errors.New("tsyms argument must not be nil")
	} else if len(tsyms) == 0 {
		return nil, errors.New("tsyms argument must not be empty")
	}

	fsymsJoinedByComma := strings.Join(fsyms, ",")
	tsymsJoinedByComma := strings.Join(tsyms, ",")

	queryParameters := map[string]string{
		"fsyms": fsymsJoinedByComma,
		"tsyms": tsymsJoinedByComma,
	}

	resp := interfaceService.CryptoComparePriceMultifullResponse{}
	err := service.DoRequest(http.MethodGet, PriceMultifullAction, queryParameters, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func (service *cryptoCompareService) DoRequest(method, action string, queryParameters map[string]string, body, writeRespToObject interface{}) error {
	var err error
	timeout := time.Duration(30 * time.Second)
	httpClient := http.Client{
		Timeout: timeout,
	}

	var httpRequest *http.Request
	httpRequest, err = service.MakeRequest(method, action, queryParameters, body)
	if err != nil {
		return err
	}
	requestUrl := httpRequest.URL.String()
	log.Println(requestUrl, method, "Trying to send request")
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	log.Println(requestUrl, method, "Request sent and returned status code ", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprint("cryptocompare returned bad http response status:", resp.Status))
	}

	errorObject := interfaceService.CryptoCompareErrorResponse{}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(respBody, &errorObject)
	if errorObject.Response == "Error" {
		stringResponse := string(respBody)
		return errors.New(fmt.Sprint("cryptocompare returned error response - ", stringResponse))
	}

	if writeRespToObject != nil {
		err = json.Unmarshal(respBody, writeRespToObject)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *cryptoCompareService) MakeRequest(method, action string, queryParameters map[string]string, body interface{}) (*http.Request, error) {

	var bodyReader io.Reader
	if body != nil {
		serializedBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewBuffer(serializedBody)
	}

	request, err := http.NewRequest(method, fmt.Sprint(service.baseAddress, "/", action), bodyReader)
	if err != nil {
		return nil, err
	}

	if bodyReader != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	if queryParameters != nil && len(queryParameters) > 0 {
		q := request.URL.Query()
		for key, value := range queryParameters {
			q.Add(key, value)
		}

		request.URL.RawQuery = q.Encode()
	}
	return request, err
}

func NewCryptocompareService(baseUrl string) *cryptoCompareService {
	if baseUrl == "" {
		panic("baseUrl argument should not be empty")
	}

	service := new(cryptoCompareService)
	service.baseAddress = baseUrl
	return service
}
