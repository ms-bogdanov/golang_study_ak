package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExchanger struct {
	mock.Mock
}

func (m *MockExchanger) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	args := m.Called(pair, limit, start, end)
	return args.Get(0).(CandlesHistory), args.Error(1)
}

func TestGeneralIndicator_GetData_Success(t *testing.T) {
	mockExchanger := new(MockExchanger)
	mockCandles := CandlesHistory{
		Candles: []Candle{
			{T: 1620000000, O: 50000, C: 51000, H: 52000, L: 49000, V: 100},
			{T: 1620003600, O: 51000, C: 51500, H: 52500, L: 50500, V: 200},
		},
	}
	mockExchanger.On("GetCandlesHistory", "BTC_USD", 30, mock.Anything, mock.Anything).Return(mockCandles, nil)

	indicatorSMA := NewIndicatorSMA(mockExchanger)
	generalIndicator := &GeneralIndicator{}
	from := time.Now().Add(-time.Hour * 24 * 5)
	to := time.Now()

	data, err := generalIndicator.GetData("BTC_USD", 30, from, to, indicatorSMA)

	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestGeneralIndicator_GetData_EmptyHistory(t *testing.T) {
	mockExchanger := new(MockExchanger)
	mockCandles := CandlesHistory{
		Candles: []Candle{},
	}
	mockExchanger.On("GetCandlesHistory", "BTC_USD", 30, mock.Anything, mock.Anything).Return(mockCandles, nil)

	indicatorSMA := NewIndicatorSMA(mockExchanger)
	generalIndicator := &GeneralIndicator{}
	from := time.Now().Add(-time.Hour * 24 * 5)
	to := time.Now()

	data, err := generalIndicator.GetData("BTC_USD", 30, from, to, indicatorSMA)

	assert.NoError(t, err)
	assert.Empty(t, data)
}

func TestGetCandlesHistory_Success(t *testing.T) {
	e := &Exmo{
		url:    "https://api.exmo.com",
		client: &http.Client{},
	}

	pair := "BTC_USD"
	limit := 60
	start := time.Unix(1609459200, 0)
	end := time.Unix(1609545600, 0)

	httpmock.ActivateNonDefault(e.client)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{"candles":[{"t":1609459200,"o":29000,"c":29500,"h":30000,"l":28500,"v":1000}]}`
	httpmock.RegisterResponder("GET", "https://api.exmo.com/candles_history?symbol=BTC_USD&resolution=60&from=1609459200&to=1609545600",
		httpmock.NewStringResponder(200, mockResponse))

	history, err := e.GetCandlesHistory(pair, limit, start, end)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(history.Candles) != 1 {
		t.Fatalf("expected 1 candle, got %d", len(history.Candles))
	}

	if history.Candles[0].O != 29000 {
		t.Errorf("expected open price 29000, got %f", history.Candles[0].O)
	}
}

func TestCalculateSMAFromCandleHistory(t *testing.T) {
	mockExchanger := new(MockExchanger)
	mockCandles := CandlesHistory{
		Candles: []Candle{
			{T: 1620000000, O: 50000, C: 51000, H: 52000, L: 49000, V: 100},
			{T: 1620003600, O: 51000, C: 51500, H: 52500, L: 50500, V: 200},
		},
	}
	mockExchanger.On("GetCandlesHistory", "BTC_USD", 30, mock.Anything, mock.Anything).Return(mockCandles, nil)

	indicatorSMA := &IndicatorSMA{exmo: mockExchanger}
	from := time.Now().Add(-time.Hour * 24 * 5)
	to := time.Now()
	result, err := indicatorSMA.GetData("BTC_USD", 30, 5, from, to)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
}

func TestCorrectlyCalculatesEMA(t *testing.T) {
	mockExchanger := new(MockExchanger)
	mockCandles := CandlesHistory{
		Candles: []Candle{
			{C: 1.0},
			{C: 2.0},
			{C: 3.0},
			{C: 4.0},
			{C: 5.0},
		},
	}
	mockExchanger.On("GetCandlesHistory", "BTC_USD", 30, mock.Anything, mock.Anything).Return(mockCandles, nil)

	indicatorEMA := NewIndicatorEMA(mockExchanger)
	result, err := indicatorEMA.GetData("BTC_USD", 30, 5, time.Now().Add(-time.Hour*24*5), time.Now())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 5, len(result))
}
