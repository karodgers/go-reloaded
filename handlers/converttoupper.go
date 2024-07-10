package handlers

// converting instances of (Up) to Upper
func ConvertUpToUpper(word string) string {
	var upWord string
	for _, value := range word {
		if value >= 'a' && value <= 'z' {
			upWord = upWord + string(value-32)
		} else {
			upWord = upWord + string(value)
		}
	}
	return upWord
}
