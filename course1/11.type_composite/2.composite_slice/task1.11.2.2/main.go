package main

func MaxDifference(numbers []int) int {

	if len(numbers) <= 1 {
		return 0
	}

	max := numbers[0]
	min := numbers[0]

	for _, v := range numbers {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return max - min
}
