package main

func RemoveExtraMemory(xs []int) []int {
	res := make([]int, len(xs))
	copy(res, xs)
	return res
}
