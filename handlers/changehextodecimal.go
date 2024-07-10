package handlers

import "fmt"

// converting Hex values to decimal values
func ChangeHexToDecimal(hexStr string) string {
	decimal := 0
	base := 1
	for i := len(hexStr) - 1; i >= 0; i-- {
		digit := hexStr[i]
		if digit >= '0' && digit <= '9' {
			decimal += int(digit-'0') * base
		} else if digit >= 'A' && digit <= 'F' {
			decimal += int(digit-'A'+10) * base
		} else if digit >= 'a' && digit <= 'f' {
			decimal += int(digit-'a'+10) * base
		}
		base *= 16
	}
	return fmt.Sprintf("%d", decimal)
}
