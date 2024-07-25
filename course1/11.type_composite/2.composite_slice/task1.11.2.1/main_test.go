package main

import (
	"reflect"
	"testing"
)

func TestGetSubSlice(t *testing.T) {

	type TestSubSlice struct {
		slice []int
		start int
		end   int
		ex    []int
	}

	subSlice := []TestSubSlice{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 6, []int{3, 4, 5, 6}},
	}

	for _, tc := range subSlice {
		res := getSubSlice(tc.slice, tc.start, tc.end)
		if reflect.DeepEqual(res, tc.ex) == false {
			t.Errorf("getSubSlice(%v, %v, %v) got %v expected %v", tc.slice, tc.start, tc.end, res, tc.ex)
		}
	}
}
