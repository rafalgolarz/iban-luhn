package cards

/*
 * MIT Licence
 * author: rafalgolarz.com
 */

import (
	"strconv"
	"strings"
)

//IsLuhn returns true if number passes the Luhn validation
func IsLuhn(number string) bool {

	//do some clean ups on white spaces
	numberFields := strings.Fields(number)
	number = strings.Join(numberFields, "")

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
