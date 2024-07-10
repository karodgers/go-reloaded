package handlers

import (
	"testing"

	"go-reloaded/handlers"
)

func TestConvertLowToLower(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "Apple", expected: "apple"},
		{input: "Banana", expected: "banana"},
		{input: "OrAnGe", expected: "orange"},
		{input: "KiWi", expected: "kiwi"},
		{input: "", expected: ""},
	}

	for _, tc := range testCases {
		result := handlers.ConvertLowToLower(tc.input)
		if result != tc.expected {
			t.Errorf("Expected ConvertLowToLower(%s) to be %s, but got %s", tc.input, tc.expected, result)
		}
	}
}
