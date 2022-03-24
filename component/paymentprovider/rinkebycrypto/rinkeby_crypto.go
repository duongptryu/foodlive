package rinkebycrypto

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type cryptoPayment struct {
	ApiKey string
}

type RespPrice struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  struct {
		EthUsd string `json:"ethusd"`
	} `json:"result"`
}

type RespTxn struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

func NewCryptoPayment(apiKey string) *cryptoPayment {
	return &cryptoPayment{
		ApiKey: apiKey,
	}
}

func (c *cryptoPayment) ParsePriceToEth(ctx context.Context, priceDola float64) (float64, error) {
	url := "https://api-rinkeby.etherscan.io/api?module=stats&action=ethprice&apikey=" + c.ApiKey
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return 0, err
	}
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	var result RespPrice
	if err = json.Unmarshal(body, &result); err != nil {
		return 0, nil
	}

	priceEth, err := strconv.ParseFloat(result.Result.EthUsd, 64)
	if err != nil {
		return 0, err
	}

	realPrice := priceDola / priceEth

	return realPrice, nil
}

func (c *cryptoPayment) CheckStatusTxn(ctx context.Context, txnHash string) (string, error) {
	url := fmt.Sprintf("https://api-rinkeby.etherscan.io/api?module=transaction&action=gettxreceiptstatus&txhash=%v&apikey=%v", txnHash, c.ApiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var result RespTxn
	if err = json.Unmarshal(body, &result); err != nil {
		return "", nil
	}

	return result.Result.Status, nil
}
