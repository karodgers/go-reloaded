package handlers

import "strings"

// converting a/A instances to an/An
func AToAn(word []string) []string {
	vowels := "aeiouhAEIOUH"
	for i := 0; i < len(word); i++ {
		if word[i] == "a" && strings.ContainsRune(vowels, rune(word[i+1][0])) {
			word[i] = "an"
		} else if word[i] == "A" && strings.ContainsRune(vowels, rune(word[i+1][0])) {
			word[i] = "An"
		}
	}
	return word
}
