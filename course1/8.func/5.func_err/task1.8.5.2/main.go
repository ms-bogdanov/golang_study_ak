package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(CheckDiscount(3000, 49))
}

func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 50 {
		return 0, errors.New("cкидка не может превышать 50%")
	}

	return price * (100 - discount) / 100, nil
}
