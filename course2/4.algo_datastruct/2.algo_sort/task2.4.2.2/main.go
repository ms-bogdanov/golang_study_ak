package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j -= 1
		}

		arr[j+1] = key
	}
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	answer := make([]int, 0, len(arr))
	less := []int{}
	more := []int{}
	pivot := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > pivot {
			more = append(more, arr[i])
		} else {
			less = append(less, arr[i])
		}
	}

	answer = append(answer, quickSort(less)...)
	answer = append(answer, pivot)
	answer = append(answer, quickSort(more)...)
	return answer
}

func GeneralSort(arr []int) {
	if len(arr) < 10 {
		insertionSort(arr)
	} else {
		selectionSort(arr)
	}
}

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original: ", data)
	sortedData := mergeSort(data)
	fmt.Println("Sorted by Merge Sort: ", sortedData)
	data = []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(data)
	fmt.Println("Sorted by Insertion Sort: ", data)
	data = []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(data)
	fmt.Println("Sorted by Selection Sort: ", data)
	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = quickSort(data)
	fmt.Println("Sorted by Quicksort: ", sortedData)
	data = []int{64, 34, 25, 12, 22, 11, 90, 55, 33, 55, 22, 11, 66}
	GeneralSort(data)
	fmt.Println("Sorted by GeneralSort: ", data)
}
