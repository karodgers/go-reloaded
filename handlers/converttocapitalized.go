package handlers

// converting instances of (Cap) to Capitaized
func ConvertCapToCapitalized(word string) string {
	if len(word) == 0 {
		return word
	}
	firstChar := word[0]
	if firstChar >= 'a' && firstChar <= 'z' {
		firstChar = firstChar - 32 // Convert to uppercase
	}
	return string(firstChar) + word[1:]
}
