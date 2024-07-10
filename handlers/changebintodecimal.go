package handlers

import "fmt"

// converting Bin values to decimal values
func ChangeBinToDecimal(binStr string) string {
	decimal := 0
	base := 1
	for i := len(binStr) - 1; i >= 0; i-- {
		if binStr[i] == '1' {
			decimal += base
		}
		base *= 2
	}
	return fmt.Sprintf("%d", decimal)
}
