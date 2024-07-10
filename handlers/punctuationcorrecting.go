package handlers

// correcting the first and second apostrophe
func PunctuationCorrecting(words []string) []string {
	count := 0
	for i, word := range words {
		if word == "'" && count == 0 {
			count += 1
			words[i+1] = word + words[i+1]
			words = append(words[:i], words[i+1:]...)
		}
	}

	for i, word := range words {
		if word == "'" {
			words[i-1] = words[i-1] + word
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
