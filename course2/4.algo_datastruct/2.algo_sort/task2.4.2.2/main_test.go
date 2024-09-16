package main

import (
	"reflect"
	"testing"
)

func TestMergeSortPositiveIntegers(t *testing.T) {
	input := []int{5, 3, 8, 6, 2, 7, 4, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := mergeSort(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMergeSortNegativeIntegers(t *testing.T) {
	input := []int{-5, 3, 8, 6, 2, -7, 4, 1}
	expected := []int{-7, -5, 1, 2, 3, 4, 6, 8}
	result := mergeSort(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{4, 2, 3, 1, 5}
	expected := []int{1, 2, 3, 4, 5}
	insertionSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, but got %v", expected, arr)
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	selectionSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, but got %v", expected, arr)
	}
}

func TestQuickSort(t *testing.T) {
	input := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	result := quickSort(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestGeneralSort(t *testing.T) {
	arr := []int{5, 2, 9, 1, 5, 6}
	expected := []int{1, 2, 5, 5, 6, 9}
	GeneralSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, but got %v", expected, arr)
	}
}
