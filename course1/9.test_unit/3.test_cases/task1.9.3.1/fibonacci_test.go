package main

import "testing"

func TestFibonacci(t *testing.T) {
	type TestData struct {
		input    int
		expected int
	}

	testCases := []TestData{
		{input: 5, expected: 5},
		{input: 6, expected: 8},
		{input: 7, expected: 13},
		{input: 8, expected: 21},
		{input: 9, expected: 34},
		{input: 10, expected: 55},
	}

	for _, tc := range testCases {
		res := Fibonacci(tc.input)
		if res != tc.expected {
			t.Errorf("Fibonacci(%d) = %d; wanted: %d", tc.input, res, tc.expected)
		}
	}
}
