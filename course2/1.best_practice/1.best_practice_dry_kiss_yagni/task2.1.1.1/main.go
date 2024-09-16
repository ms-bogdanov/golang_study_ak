package main

import "fmt"

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

func NewStatisticProfit(opts ...func(*StatisticProfit)) Profitable {
	sp := &StatisticProfit{}
	for _, opt := range opts {
		opt(sp)
	}
	return sp
}

func WithAverageProfit(s *StatisticProfit) {
	s.getAverageProfit = func() float64 {
		return s.Average(s.product.Sells) - s.Average(s.product.Buys)
	}
}

func WithAverageProfitPercent(s *StatisticProfit) {
	s.getAverageProfitPercent = func() float64 {
		averageProfit := s.getAverageProfit()
		averageBuys := s.Average(s.product.Buys)
		if averageBuys == 0 {
			return 0
		}
		return (averageProfit / averageBuys) * 100
	}
}

func WithCurrentProfit(s *StatisticProfit) {
	s.getCurrentProfit = func() float64 {
		return s.product.CurrentPrice - s.product.CurrentPrice*(100-s.product.ProfitPercent)/100
	}
}

func WithDifferenceProfit(s *StatisticProfit) {
	s.getDifferenceProfit = func() float64 {
		return s.product.CurrentPrice - s.Average(s.product.Sells)
	}
}

func WithAllData(s *StatisticProfit) {
	s.getAllData = func() []float64 {
		res := make([]float64, 0, 4)
		if s.getAverageProfitPercent != nil {
			res = append(res, s.getAverageProfitPercent())
		}
		if s.getCurrentProfit != nil {
			res = append(res, s.getCurrentProfit())
		}
		if s.getDifferenceProfit != nil {
			res = append(res, s.getDifferenceProfit())
		}
		return res
	}
}

func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}

func (s *StatisticProfit) GetAverageProfit() float64 {
	return s.getAverageProfit()
}

func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	return s.getAverageProfitPercent()
}

func (s *StatisticProfit) GetCurrentProfit() float64 {
	return s.getCurrentProfit()
}

func (s *StatisticProfit) GetDifferenceProfit() float64 {
	return s.getDifferenceProfit()
}

func (s *StatisticProfit) GetAllData() []float64 {
	return s.getAllData()
}

func (s *StatisticProfit) Average(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}
	return s.Sum(prices) / float64(len(prices))
}

func (s *StatisticProfit) Sum(prices []float64) float64 {
	var res float64
	for _, price := range prices {
		res += price
	}
	return res
}

func main() {
	product := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}

	statProfit := NewStatisticProfit(
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDifferenceProfit,
		WithAllData,
	).(*StatisticProfit)

	statProfit.SetProduct(product)
	fmt.Println("Average Profit:", statProfit.GetAverageProfit())
	fmt.Println("Average Profit Percent:", statProfit.GetAverageProfitPercent())
	fmt.Println("Current Profit:", statProfit.GetCurrentProfit())
	fmt.Println("Difference Profit:", statProfit.GetDifferenceProfit())
	fmt.Println("All Data:", statProfit.GetAllData())
}

// type CocaCola struct {
// 	Sells         []float64
// 	Buys          []float64
// 	CurrentPrice  float64
// 	ProfitPercent float64
// }

// type Pepsi struct {
// 	Sells         []float64
// 	Buys          []float64
// 	CurrentPrice  float64
// 	ProfitPercent float64
// }

// func getAverageProfitCocaCola(c CocaCola) float64 {
// 	return 0
// }

// func getAverageProfitPepsi(c Pepsi) float64 {
// 	return 0
// }

// func getCurrentProfitCocaCola(c CocaCola) float64 {
// 	return 0
// }

// func getCurrentProfitPepsi(c Pepsi) float64 {
// 	return 0
// }

// func getDifferenceProfitCocaCola(c CocaCola) float64 {
// 	return 0
// }

// func getDifferenceProfitPepsi(c Pepsi) float64 {
// 	return 0
// }
