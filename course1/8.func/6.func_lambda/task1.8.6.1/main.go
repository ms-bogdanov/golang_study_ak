package main

import "fmt"

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))
	fmt.Println(MathOperate(Mul, 1, 7, 3))
	fmt.Println(MathOperate(Sub, 13, 2, 3))

}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}

func Sum(a ...int) int {
	res := 0
	for _, n := range a {
		res += n
	}
	return res
}
func Mul(a ...int) int {
	res := 1
	for _, n := range a {
		res *= n
	}
	return res
}
func Sub(a ...int) int {
	if len(a) == 0 {
		return 0
	}

	res := a[0]
	for i := 1; i < len(a); i++ {
		res -= a[i]
	}
	return res
}
