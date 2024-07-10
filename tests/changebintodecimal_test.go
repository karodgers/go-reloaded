package handlers

import (
	"testing"

	"go-reloaded/handlers"
)

func TestChangeBinToDecimal(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "1101", expected: "13"},
		{input: "1010", expected: "10"},
		{input: "11111111", expected: "255"},
		{input: "10000", expected: "16"},
	}

	for _, tc := range testCases {
		result := handlers.ChangeBinToDecimal(tc.input)
		if result != tc.expected {
			t.Errorf("Expected ChangeBinToDecimal(%s) to be %s, but got %s", tc.input, tc.expected, result)
		}
	}
}
