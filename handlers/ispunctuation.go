package handlers

import "unicode"

// checking instances of punctuation marks

func IsPunctuation(word string) bool {
	if len(word) >= 0 || len(word) <= 9 {
		return unicode.IsPunct(rune(word[0]))
	}
	return false
}
