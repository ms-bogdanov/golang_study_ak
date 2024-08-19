package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTickerReturnsValidData(t *testing.T) {
	mockResponse := `{
    "BTC_USD": {
      "buy_price": "50000",
      "sell_price": "51000",
      "last_trade": "50500",
      "high": "52000",
      "low": "48000",
      "avg": "50000",
      "vol": "1000",
      "vol_curr": "50000000",
      "updated": 1617184800
    }
  }`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	exmo := NewExmo(WithURL(server.URL))

	ticker, err := exmo.GetTicker()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := Ticker{
		"BTC_USD": {
			BuyPrice:  "50000",
			SellPrice: "51000",
			LastTrade: "50500",
			High:      "52000",
			Low:       "48000",
			Avg:       "50000",
			Vol:       "1000",
			VolCurr:   "50000000",
			Updated:   1617184800,
		},
	}

	if !jsonEqual(ticker, expected) {
		t.Errorf("expected %v, got %v", expected, ticker)
	}
}

func TestGetTrades_SuccessfullyFetchTradesForSinglePair(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string][]Pair{
			"BTC_USD": {
				{TradeID: 1, Date: time.Now().Unix(), Type: "buy", Quantity: "0.1", Price: "50000", Amount: "5000"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	exmo := NewExmo(WithURL(server.URL))
	trades, err := exmo.GetTrades("BTC_USD")

	assert.NoError(t, err)
	assert.NotNil(t, trades)
	assert.Contains(t, trades, "BTC_USD")
	assert.Equal(t, 1, len(trades["BTC_USD"]))
}

func jsonEqual(a, b interface{}) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}

func TestGetOrderBookSinglePair(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]OrderBookPair{
			"BTC_USD": {
				AskQuantity: "10",
				AskAmount:   "100",
				AskTop:      "10",
				BidQuantity: "5",
				BidAmount:   "50",
				BidTop:      "5",
				Ask:         [][]string{{"10", "1"}},
				Bid:         [][]string{{"5", "1"}},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	exmo := NewExmo(WithURL(server.URL))

	orderBook, err := exmo.GetOrderBook(10, "BTC_USD")

	assert.NoError(t, err)
	assert.NotNil(t, orderBook)
	assert.Equal(t, "10", orderBook["BTC_USD"].AskQuantity)
}

func TestGetCurrencies_Success(t *testing.T) {
	mockResponse := `["USD", "EUR", "BTC"]`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	exmo := NewExmo(WithURL(server.URL))
	currencies, err := exmo.GetCurrencies()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := Currencies{"USD", "EUR", "BTC"}
	if len(currencies) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, currencies)
	}

	for i, currency := range currencies {
		if currency != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], currency)
		}
	}
}

func TestGetClosePriceReturnsCorrectPrices(t *testing.T) {
	mockResponse := CandlesHistory{
		Candles: []Candle{
			{C: 100.0},
			{C: 200.0},
			{C: 300.0},
		},
	}
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer mockServer.Close()

	exmo := NewExmo(WithURL(mockServer.URL))
	pair := "BTC_USD"
	limit := 1
	start := time.Now().Add(-24 * time.Hour)
	end := time.Now()

	prices, err := exmo.GetClosePrice(pair, limit, start, end)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedPrices := []float64{100.0, 200.0, 300.0}

	for i, price := range prices {
		if price != expectedPrices[i] {
			t.Errorf("expected price %v, got %v", expectedPrices[i], price)
		}
	}
}

func TestIndicator(t *testing.T) {
	mockResponse := CandlesHistory{
		Candles: []Candle{
			{C: 100.0},
			{C: 200.0},
			{C: 300.0},
			{C: 400.0},
			{C: 500.0},
		},
	}
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer mockServer.Close()

	exchange := NewExmo(WithURL(mockServer.URL))
	ind := NewIndicator(exchange, WithCalculateEMA(calculateEMA), WithCalculateSMA(calculateSMA))
	result, err := ind.SMA("BTC_USD", 5, 3, time.Now().AddDate(0, 0, -2), time.Now())

	if err != nil {
		t.Errorf("got unexpected error: %v", err)
	}

	expected := []float64{100.0, 150.0, 200.0, 300.0, 400.0}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	result, err = ind.EMA("BTC_USD", 5, 3, time.Now().AddDate(0, 0, -2), time.Now())
	expected = []float64{100.0, 150.0, 225.0, 312.5, 406.25}
	assert.NoError(t, err)
	assert.Equal(t, result, expected)
}
