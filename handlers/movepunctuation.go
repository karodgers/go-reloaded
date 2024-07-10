package handlers

// correcting punctuation spacing
func MovePunctuation(slice []string) []string {
	var result []string

	for i := 0; i < len(slice); i++ {
		word := slice[i]
		if i < len(slice)-1 && IsPunctuation(slice[i+1]) {
			word += slice[i+1]
			i++ // Skip the next word (punctuation)
		}
		result = append(result, word)
	}

	return result
}
