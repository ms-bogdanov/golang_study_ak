package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func DecimalSum(a, b string) (res string, err error) {
	var aDecimal, bDecimal decimal.Decimal

	aDecimal, err = decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	bDecimal, err = decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	resDecimal := decimal.Sum(aDecimal, bDecimal)

	return resDecimal.String(), nil
}

func DecimalSubstract(a, b string) (res string, err error) {
	var aDecimal, bDecimal decimal.Decimal

	aDecimal, err = decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	bDecimal, err = decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	resDecimal := aDecimal.Sub(bDecimal)

	return resDecimal.String(), nil
}

func DecimalMultiply(a, b string) (res string, err error) {
	var aDecimal, bDecimal decimal.Decimal

	aDecimal, err = decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	bDecimal, err = decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	resDecimal := aDecimal.Mul(bDecimal)

	return resDecimal.String(), nil
}

func DecimalDivide(a, b string) (res string, err error) {
	var aDecimal, bDecimal decimal.Decimal

	aDecimal, err = decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	bDecimal, err = decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	resDecimal := aDecimal.Div(bDecimal)

	return resDecimal.String(), nil
}

func DecimalRound(a string, precision int32) (res string, err error) {
	var aDecimal decimal.Decimal

	aDecimal, err = decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	resDecimal := aDecimal.Round(precision)

	return resDecimal.String(), nil
}

func DecimalGreaterThan(a, b string) (res bool, err error) {
	var aDecimal, bDecimal decimal.Decimal

	aDecimal, err = decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	bDecimal, err = decimal.NewFromString(b)
	if err != nil {
		return false, err
	}

	res = aDecimal.GreaterThan(bDecimal)

	return res, nil
}

func DecimalLessThan(a, b string) (res bool, err error) {
	var aDecimal, bDecimal decimal.Decimal

	aDecimal, err = decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	bDecimal, err = decimal.NewFromString(b)
	if err != nil {
		return false, err
	}

	res = aDecimal.LessThan(bDecimal)

	return res, nil
}

func DecimalEqual(a, b string) (res bool, err error) {
	var aDecimal, bDecimal decimal.Decimal

	aDecimal, err = decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	bDecimal, err = decimal.NewFromString(b)
	if err != nil {
		return false, err
	}

	res = aDecimal.Equal(bDecimal)

	return res, nil
}

func main() {
	res, err := DecimalSum("123", "234")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	res, err = DecimalSubstract("123", "234")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	res, err = DecimalMultiply("123", "234")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	res, err = DecimalDivide("234", "123")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	res, err = DecimalRound("234.5678", 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	var resBool bool
	resBool, err = DecimalGreaterThan("234.5678", "234.1111")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resBool)
	resBool, err = DecimalLessThan("212.123456", "212.1111")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resBool)
	resBool, err = DecimalEqual("123.123456", "123.123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resBool)
}
