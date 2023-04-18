package symbol

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type coinData struct {
	MarketData marketData `json:"market_data"`
	Symbol     string     `json:"symbol"`
}

type marketData struct {
	CurrentPrice currentPrice `json:"current_price"`
	MarketCap    marketCap    `json:"market_cap"`
	TotalVolume  totalVolume  `json:"total_volume"`
}

type currentPrice struct {
	Usd float64 `json:"usd"`
}

type marketCap struct {
	Usd float64 `json:"usd"`
}

type totalVolume struct {
	Usd float64 `json:"usd"`
}

type SymbolResponse struct {
	DailyVolume float64   `json:"dailyVolume"`
	MarketCap   float64   `json:"marketCap"`
	Price       float64   `json:"price"`
	Symbol      string    `json:"symbol"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type symbolData struct {
	MarketCap float64
	Price     float64
	Symbol    string
	Volume24h float64
}

const coinGeckoAPIUrl = "https://api.coingecko.com/api/v3/coins/"

func Fetch(symbol string) (*SymbolResponse, error) {
	symbolData, err := fetchSymbolPriceAndVolume(symbol)

	if err != nil {
		return nil, err
	}

	symbolResponse := &SymbolResponse{
		DailyVolume: symbolData.Volume24h,
		MarketCap:   symbolData.MarketCap,
		Price:       symbolData.Price,
		Symbol:      symbolData.Symbol,
		UpdatedAt:   time.Now(),
	}

	return symbolResponse, nil
}

func fetchSymbolPriceAndVolume(s string) (*symbolData, error) {
	response, err := http.Get(coinGeckoAPIUrl + s)

	if err != nil {
		fmt.Println(err)

		return nil, errors.New("An error occurred.")
	}

	defer response.Body.Close()

	var data coinData

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		fmt.Println(err)

		return nil, errors.New("An error occurred.")
	}

	symbolData := &symbolData{
		Price:     data.MarketData.CurrentPrice.Usd,
		MarketCap: data.MarketData.MarketCap.Usd,
		Symbol:    data.Symbol,
		Volume24h: data.MarketData.TotalVolume.Usd,
	}

	return symbolData, nil
}
