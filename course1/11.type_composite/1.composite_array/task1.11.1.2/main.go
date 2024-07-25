package main

import (
	"sort"
)

func sortDescInt(arr [8]int) [8]int {
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr
}

func sortAscInt(arr [8]int) [8]int {
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func sortDescFloat(arr [8]float64) [8]float64 {
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr
}

func sortAscFloat(arr [8]float64) [8]float64 {
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}
