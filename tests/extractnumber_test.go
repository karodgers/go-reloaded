package handlers

import (
	"testing"

	"go-reloaded/handlers"
)

func TestExtractNumber(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "10Rogy", expected: 10},
		{input: "9f", expected: 9},
		{input: "98baibe", expected: 98},
	}

	for _, tc := range testCases {
		result := handlers.ExtractNumber(tc.input)
		if result != tc.expected {
			t.Errorf("Expected ExtractNumber(%s) to be %d, but got %d", tc.input, tc.expected, result)
		}
	}
}
