package main

import "fmt"

func main() {
	fmt.Println(Operate(Concat, "Hello, ", "World!"))
	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))
	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5))
}

func Operate(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}

func Concat(xs ...interface{}) interface{} {
	str := ""
	for _, v := range xs {
		str += v.(string)
	}
	return str
}

func Sum(xs ...interface{}) interface{} {
	switch xs[0].(type) {
	case int:
		intSum := 0
		for _, v := range xs {
			intSum += v.(int)
		}
		return intSum
	case float64:
		floatSum := 0.0
		for _, v := range xs {
			floatSum += v.(float64)
		}
		return floatSum
	}
	return nil
}
