package handlers

// converting instances of (Low) to Lower
func ConvertLowToLower(word string) string {
	var lowWord string
	for _, value := range word {
		if value >= 'A' && value <= 'Z' {
			lowWord = lowWord + string(value+32)
		} else {
			lowWord = lowWord + string(value)
		}
	}
	return lowWord
}
