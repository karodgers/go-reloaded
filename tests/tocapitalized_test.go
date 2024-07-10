package handlers

import (
	"testing"

	"go-reloaded/handlers"
)

func TestConvertCapToCapitalized(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "apple", expected: "Apple"},
		{input: "Banana", expected: "Banana"},
		{input: "orange", expected: "Orange"},
		{input: "kiwi", expected: "Kiwi"},
		{input: "", expected: ""},
	}

	for _, tc := range testCases {
		result := handlers.ConvertCapToCapitalized(tc.input)
		if result != tc.expected {
			t.Errorf("Expected ConvertCapToCapitalized(%s) to be %s, but got %s", tc.input, tc.expected, result)
		}
	}
}
