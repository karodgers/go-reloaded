package handlers

// extracting number from (cap, <number>) et al.
func ExtractNumber(s string) int {
	num := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		} else {
			break
		}
	}
	return num
}
