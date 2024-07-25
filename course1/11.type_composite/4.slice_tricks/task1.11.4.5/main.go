package main

func FilterDividers(xs []int, divider int) []int {
	res := make([]int, len(xs))
	for _, v := range xs {
		if v%divider == 0 {
			res = append(res, v)
		}
	}
	return res
}
