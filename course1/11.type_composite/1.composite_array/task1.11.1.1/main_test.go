package main

import "testing"

type testOpSum struct {
	input    [8]int
	expected int
}

func TestSum(t *testing.T) {
	testCases := []testOpSum{
		{input: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, expected: 36},
	}

	for _, tc := range testCases {
		result := sum(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}

type testOpAverage struct {
	input    [8]int
	expected float64
}

func TestAverage(t *testing.T) {
	testCases := []testOpAverage{
		{input: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, expected: 4.5},
	}

	for _, tc := range testCases {
		result := average(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}

type testOpAverageFloat struct {
	input    [8]float64
	expected float64
}

func TestAverageFloat(t *testing.T) {
	testCases := []testOpAverageFloat{
		{input: [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}, expected: 5},
	}

	for _, tc := range testCases {
		result := averageFloat(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}

type testOpReverse struct {
	input    [8]int
	expected [8]int
}

func TestReverse(t *testing.T) {
	testCases := []testOpReverse{
		{input: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, expected: [8]int{8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, tc := range testCases {
		result := reverse(tc.input)
		for i, val := range tc.expected {
			if result[i] != val {
				t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
			}
		}
	}
}
