package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/cinar/indicator"
)

type Exchanger interface {
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
}

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type Indicator interface {
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(rsiPeriod int) ([]float64, []float64)
	SMA(period int) []float64
	MACD() ([]float64, []float64)
	EMA() []float64
}

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	exmo := &Exmo{client: &http.Client{}, url: "https://api.exmo.com/v1.1"}

	for _, option := range opts {
		option(exmo)
	}

	return exmo
}

type GeneralIndicatorer interface {
	GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64,
		error)
}

type Indicatorer interface {
	GetData(period int, history CandlesHistory) ([]float64, error)
}

type IndicatorSMA struct {
	candleHistory CandlesHistory
	exmo          Exchanger
}

type IndicatorEMA struct {
	candleHistory CandlesHistory
	exmo          Exchanger
}

func (i *IndicatorSMA) GetData(period int, history CandlesHistory) ([]float64, error) {
	i.candleHistory = history
	return i.SMA(period), nil
}

func (t *IndicatorSMA) SMA(period int) []float64 {
	closing := []float64{}

	for _, i := range t.candleHistory.Candles {
		closing = append(closing, i.C)
	}

	return indicator.Sma(period, closing)
}

func (i *IndicatorEMA) GetData(period int, history CandlesHistory) ([]float64, error) {
	i.candleHistory = history
	return i.EMA(period), nil
}

func (i *IndicatorEMA) EMA(period int) []float64 {
	closing := []float64{}

	for _, j := range i.candleHistory.Candles {
		closing = append(closing, j.C)
	}

	return indicator.Ema(period, closing)
}

func GetRespBody(e *Exmo, request string) ([]byte, error) {
	resp, err := e.client.Get(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func UnmarshalCandles(data []byte) (CandlesHistory, error) {
	var r CandlesHistory
	err := json.Unmarshal(data, &r)
	return r, err
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	request := e.url + "/candles_history?symbol=" + pair + "&resolution=" + strconv.Itoa(limit) + "&from=" +
		strconv.FormatInt(start.Unix(), 10) + "&to=" + strconv.FormatInt(end.Unix(), 10)
	body, err := GetRespBody(e, request)
	var history CandlesHistory

	if err != nil {
		return history, err
	}

	history, err = UnmarshalCandles(body)
	if err != nil {
		return history, err
	}

	return history, nil
}

func NewIndicatorSMA(exchange Exchanger) *IndicatorSMA {
	return &IndicatorSMA{exmo: exchange}
}

func NewIndicatorEMA(exchange Exchanger) *IndicatorEMA {
	return &IndicatorEMA{exmo: exchange}
}

type GeneralIndicator struct {
}

type Dashboarder interface {
	GetDashboard(pair string, opts ...IndicatorOpt) (DashboardData, error)
}

type DashboardData struct {
	Name           string
	CandlesHistory CandlesHistory
	Indicators     map[string][]IndicatorData
	from           time.Time
	to             time.Time
}

type IndicatorData struct {
	Name     string
	Period   int
	Indicate []float64
}

type IndicatorOpt struct {
	Name      string
	Periods   []int
	Indicator Indicatorer
}

type Dashboard struct {
	exchange           Exchanger
	withCandlesHistory bool
	IndicatorOpts      []IndicatorOpt
	period             int
	from               time.Time
	to                 time.Time
}

func NewDashboard(exchange Exchanger) *Dashboard {
	return &Dashboard{exchange: exchange}
}

func (d *Dashboard) WithCandlesHistory(period int, from, to time.Time) {
	d.period = period
	d.from = from
	d.to = to
	d.withCandlesHistory = true
}

func (d *Dashboard) GetDashboard(pair string, opts ...IndicatorOpt) (DashboardData, error) {
	data := DashboardData{Name: pair, from: d.from, to: d.to, Indicators: make(map[string][]IndicatorData)}
	candlesHistory, err := d.exchange.GetCandlesHistory(pair, d.period, d.from, d.to)

	if err != nil {
		return data, err
	}

	data.CandlesHistory = candlesHistory

	for _, i := range opts {
		indicat := []IndicatorData{}
		for _, j := range i.Periods {
			currentData := IndicatorData{Name: i.Name, Period: j}
			indicateData, err := i.Indicator.GetData(j, data.CandlesHistory)

			if err != nil {
				fmt.Println("we have error with getting data")
				break
			}

			currentData.Indicate = indicateData
			indicat = append(indicat, currentData)
		}
		data.Indicators[i.Name] = indicat

	}

	return data, nil
}

func main() {
	exchange := NewExmo()
	dashboard := NewDashboard(exchange)
	dashboard.WithCandlesHistory(30, time.Now().Add(-time.Hour*24*5), time.Now())
	opts := []IndicatorOpt{
		IndicatorOpt{
			Name:      "SMA",
			Periods:   []int{5, 10, 20},
			Indicator: NewIndicatorSMA(exchange),
		},
		IndicatorOpt{
			Name:      "EMA",
			Periods:   []int{5, 10, 20},
			Indicator: NewIndicatorEMA(exchange),
		},
	}
	data, err := dashboard.GetDashboard("BTC_USD", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
