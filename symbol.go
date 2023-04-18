package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PriceResponse struct {
	Price  float64 `json:"price,string"`
	Symbol string  `json:"symbol"`
}

type SymbolResponse struct {
	DailyVolume float64   `json:"dailyVolume"`
	Price       float64   `json:"price"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type VolumeResponse struct {
	BaseVol  float64 `json:"volume,string"`
	QuoteVol float64 `json:"quoteVolume,string"`
	Symbol   string  `json:"symbol"`
}

const baseAPIUrl = "https://api.binance.com/api/v3/ticker"
const pricePath = "price?symbol="
const volumePath = "24hr?symbol="

func Fetch(symbol string) SymbolResponse {
	price := fetchPrice(symbol)
	volume := fetchVolume(symbol)

	symbolResponse := SymbolResponse{
		DailyVolume: volume,
		Price:       price,
		UpdatedAt:   time.Now(),
	}

	return symbolResponse
}

func main() {
	Fetch("BTCUSDT")
}

func fetchPrice(symbol string) float64 {
	priceURL := fmt.Sprintf(baseAPIUrl+"/"+pricePath+"%s", symbol)

	priceResp, err := http.Get(priceURL)

	if err != nil {
		panic(err)
	}

	defer priceResp.Body.Close()

	var priceResponse PriceResponse

	err = json.NewDecoder(priceResp.Body).Decode(&priceResponse)

	if err != nil {
		panic(err)
	}

	return priceResponse.Price
}

func fetchVolume(symbol string) float64 {
	volumeURL := fmt.Sprintf(baseAPIUrl+"/"+volumePath+"%s", symbol)
	volumeResp, err := http.Get(volumeURL)

	if err != nil {
		panic(err)
	}

	defer volumeResp.Body.Close()

	var volumeResponse VolumeResponse

	err = json.NewDecoder(volumeResp.Body).Decode(&volumeResponse)

	if err != nil {
		panic(err)
	}

	return volumeResponse.QuoteVol
}
