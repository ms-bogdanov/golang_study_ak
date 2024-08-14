package main

import "testing"

type testAppendInt struct {
	slice   []int
	numbers []int
	ex      []int
}

func TestAppendInt(t *testing.T) {
	testCases := []testAppendInt{
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
	}

	for _, tc := range testCases {
		appendInt(&tc.slice, tc.numbers...)
		for i, val := range tc.ex {
			if tc.slice[i] != val {
				t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.slice, tc.ex, tc.slice)
			}
		}
	}
}
