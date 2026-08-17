package main

import "testing"

func TestCases(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{10, 90, 100},
		{1, 1, 2},
		{0, 0, 0},
		{-20, -30, -50},
		{20, -30, -10},
	}

	for _, test := range tests {
		result := sum(test.a, test.b)
		if result != test.expected {
			t.Errorf("sum(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}
