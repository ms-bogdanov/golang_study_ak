package main

import (
	"testing"
)

type TestDiff struct {
	slice []int
	ex    int
}

func TestMaxDifference(t *testing.T) {

	testCases := []TestDiff{
		{[]int{1, 5, 10, 6, 3, 8}, 9},
		{[]int{}, 0},
		{[]int{1}, 0},
	}

	for _, tc := range testCases {
		res := MaxDifference(tc.slice)
		if res != tc.ex {
			t.Errorf("Unexpected result. Slice: %v, Expected: %v, Got: %v", tc.slice, tc.ex, res)
		}
	}
}
