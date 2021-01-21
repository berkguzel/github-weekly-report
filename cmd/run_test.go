package main

import (
	"testing"
)

func TestReturnPercentage(t *testing.T) {
	got := ReturnPercentage(10, 20)
	want := 100.00

	if got != want {
		t.Errorf("got %.2f, wanted %.2f", got , want)
	}
}

func TestReturnPercentageExtensive(t *testing.T) {
	var tests = []struct {
		inputBeginning int
		inputFinishing int
		expected float64
	}{
		{10, 5, -50},
		{100, 0, -100},
		{1, 10, 900},
		{1, 2, 100},
		{0, 1, 100},
		{0, 20, 2000},
		{0, 250, 25000},
		{0, 0, 0},
	}

	for _, test := range tests {
		if result := ReturnPercentage(test.inputBeginning, test.inputFinishing);
		result != test.expected {
			t.Errorf("Test failed: beginning: %d finishing: %d, result: %.2f", 
			test.inputBeginning, test.inputFinishing, result)
		}	
	}
}

