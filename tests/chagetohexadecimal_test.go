package handlers

import (
	"testing"

	"go-reloaded/handlers"
)

func TestChangeHexToDecimal(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "1A", expected: "26"},
		{input: "FF", expected: "255"},
		{input: "5", expected: "5"},
		{input: "B4", expected: "180"},
	}

	for _, tc := range testCases {
		result := handlers.ChangeHexToDecimal(tc.input)
		if result != tc.expected {
			t.Errorf("Expected ChangeHexToDecimal(%s) to be %s, but got %s", tc.input, tc.expected, result)
		}
	}
}
