package main

import "testing"

func TestCases(t *testing.T) {
	tests := []struct {
		raio     float64
		expected string
	}{
		{1.0, "3.1416"},
		{2.0, "12.5664"},
		{100.64, "31819.3103"},
		{150.0, "70685.7750"},
	}

	for _, test := range tests {
		result := calculateArea(test.raio)
		if result != test.expected {
			t.Errorf("calculateArea(%.2f) = %s; expected %s", test.raio, result, test.expected)
		}
	}
}
