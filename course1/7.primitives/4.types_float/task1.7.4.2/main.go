package main

import (
	"fmt"
	"time"

	"github.com/mattevans/dinero"
)

func main() {
	rate := currencyPairRate("USD", "EUR", 100.01)
	fmt.Println(rate)
}

func currencyPairRate(from string, to string, amount float64) float64 {
	client := dinero.NewClient("827dec1f583440849866926eb3f16106", from, 20*time.Minute)

	rsp, err := client.Rates.Get(to)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return *rsp * amount
}
