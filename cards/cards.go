package cards

/*
 * MIT Licence
 * author: rafalgolarz.com
 */

import (
	"strconv"
)

//IsLuhn returns true is number passes the Luhn validation
func IsLuhn(number string) bool {

	numberOfDigits := len(number)

	if numberOfDigits > 0 {
		sum := 0
		alt := false

		for i := numberOfDigits - 1; i >= 0; i-- {
			currentDigit, err := strconv.Atoi(string(number[i]))
			if err == nil {
				if alt {
					if currentDigit *= 2; currentDigit > 9 {
						currentDigit -= 9
					}
				}
				sum += currentDigit
				alt = !alt
			} else {
				return false
			}
		}
		return sum%10 == 0
	}

	return false
}
