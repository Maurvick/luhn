package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	sum := 0
	// For Luhn algorithm even id length is mandatory
	cardNumber := strings.ReplaceAll(id, " ", "")
	length := len(cardNumber)

	if length < 2 {
		return false
	}
	// Iterate over the string from right to left
	for i := length - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}
		// Double every second digit from the right
		if (length - i) % 2 == 0 {
			num *= 2
			// If the result is greater than 9, subtract 9
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	
 	return sum % 10 == 0
}
