package main

import "testing"

func TestCalc(t *testing.T) {
	tests := []struct{
		input int
		want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, test := range tests {
		got := Calc(test.input)
		if got != test.want {
			t.Errorf("Calc(%d) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestCalc2(t *testing.T) {
	tests := []struct{
		input int
		want int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, test := range tests {
		got := Calc2(test.input)
		if got != test.want {
			t.Errorf("Calc2(%d) = %d, want %d", test.input, got, test.want)
		}
	}
}
