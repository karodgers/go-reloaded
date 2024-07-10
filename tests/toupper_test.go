package handlers

import (
	"testing"

	"go-reloaded/handlers"
)

func TestConvertUpToUpper(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "apple", expected: "APPLE"},
		{input: "banana", expected: "BANANA"},
		{input: "orange", expected: "ORANGE"},
		{input: "kiwi", expected: "KIWI"},
		{input: "", expected: ""},
	}

	for _, tc := range testCases {
		result := handlers.ConvertUpToUpper(tc.input)
		if result != tc.expected {
			t.Errorf("Expected ConvertUpToUpper(%s) to be %s, but got %s", tc.input, tc.expected, result)
		}
	}
}
