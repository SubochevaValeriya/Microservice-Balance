package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseCCY = "RUB"

type ResponseObject struct {
	Success bool `json:"success"`
	Query   struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int    `json:"amount"`
	} `json:"query"`
	Info struct {
		Timestamp int     `json:"timestamp"`
		Rate      float64 `json:"rate"`
	} `json:"info"`
	Historical string  `json:"historical"`
	Date       string  `json:"date"`
	Result     float64 `json:"result"`
}

func convertToCCY(ccy string, amount int) (int, error) {
	url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/convert?to=%s&from=%s&amount=%d", ccy, baseCCY, amount)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("can't make request to exchange rate API: %w", err)
	}
	req.Header.Set("apikey", os.Getenv("EXCHANGE_RATES_ID"))
	res, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("can't make request to exchange rate API: %w", err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("can't read response from exchange rate API: %w", err)
	}

	var responseObject ResponseObject

	if err = json.Unmarshal(responseData, &responseObject); err != nil {
		return 0, fmt.Errorf("can't unmarshal json response from exchange rate API: %w", err)
	}

	return int(responseObject.Result), err
}
