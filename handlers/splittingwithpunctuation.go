package handlers

import "unicode"

// split srings with punctuation marks
func SplitStringWithPunctuation(slice []string) []string {
	var result []string

	for _, str := range slice {
		var newStr string
		var prevChar rune
		var punctuationFound bool

		for _, char := range str {
			if unicode.IsPunct(char) && char != '\'' {
				// Mark that a punctuation was found
				punctuationFound = true
				newStr += string(char)
			} else {
				// If previous character was a punctuation (except single quote), add a space before appending current character
				if punctuationFound && prevChar != ' ' {
					newStr += " "
				}
				newStr += string(char)
				// Reset punctuationFound after adding the word
				punctuationFound = false
			}
			prevChar = char
		}

		result = append(result, newStr)
	}

	return result
}
