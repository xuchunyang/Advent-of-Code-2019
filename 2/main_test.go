package main

import (
	"reflect"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, test := range tests {
		got := Calc(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Calc(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
