package accounts

/*
 * MIT Licence
 * author: rafalgolarz.com
 */

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
)

/*
 * lettersToNumber letters to their numeric values
 * A = 10, B = 11, C = 12, ..., Z = 35
 */
var lettersToNumber map[string]string

/*
 * lengts of IBAN numbers (ISO 13616:2007) per country
 * first two characters of IBAN number are always two letters
 * representing a country (ISO 3166-1 alpha-2)
 * for instance, length 28 means 2 country letters + 26 digits
 */
var lengths map[string]int

//the smallest length found in the lengths map
//so we know that's the smallest account number length we accept
const minLength = 15

func init() {

	lettersToNumber = make(map[string]string)
	//let's fill the map dynamically
	for number := 10; number <= 35; number++ {
		//10 + 55 = 65 = 'A' in ASCII
		letter := string(number + 55)
		lettersToNumber[letter] = strconv.Itoa(number)
	}

	//each country may have a different length of its IBAN
	lengths = map[string]int{
		"AD": 24, "AE": 23, "AL": 28, "AT": 20, "AZ": 28, "BA": 20,
		"BE": 16, "BG": 22, "BH": 22, "BR": 29, "CH": 21, "CR": 22,
		"CY": 28, "CZ": 24, "DE": 22, "DK": 18, "DO": 28, "EE": 20,
		"ES": 24, "FI": 18, "FR": 27, "FO": 18, "GB": 22, "GE": 22,
		"GI": 23, "GL": 18, "GR": 27, "GT": 28, "HR": 21, "HU": 28,
		"IE": 22, "IL": 23, "IS": 26, "IT": 27, "JO": 30, "KW": 30,
		"KZ": 20, "LB": 28, "LI": 21, "LT": 20, "LU": 20, "LV": 21,
		"MC": 27, "MD": 24, "ME": 22, "MK": 19, "MR": 27, "MT": 31,
		"MU": 30, "NL": 18, "NO": 15, "PK": 24, "PL": 28, "PS": 29,
		"PT": 25, "QA": 29, "RO": 24, "RS": 22, "SA": 24, "SE": 24,
		"SI": 19, "SK": 24, "SM": 27, "TN": 24, "TR": 26, "VG": 24,
		//additional list used by Nordea
		"DZ": 24, "AO": 25, "BJ": 28, "BF": 27, "BI": 16, "CM": 27,
		"CV": 25, "IR": 26, "CI": 28, "MG": 27, "ML": 28, "MZ": 25,
		"SN": 28, "UA": 29,
	}

}

func getLetterValue(letter string) string {
	return lettersToNumber[letter]
}

func lettersToValues(bban string) string {
	var res string
	for i := 0; i < len(bban); i++ {
		//check if it's a letter
		if (bban[i] >= 65 && bban[i] <= 90) || (bban[i] >= 97 && bban[i] <= 122) {
			res += getLetterValue(string(bban[i]))
		} else {
			res += string(bban[i])
		}
	}
	return res
}

func getExpectedLength(twoLetters string) int {
	return lengths[twoLetters]
}

// IsIBAN returns true if accountNumber meets IBAN criteria
func IsIBAN(account string) bool {

	//do some clean ups on white spaces
	accountFields := strings.Fields(account)
	account = strings.Join(accountFields, "")

	accountLength := len(account)
	if accountLength >= minLength {
		country := account[0:2]

		if getExpectedLength(country) == accountLength {
			//BBAN - Basic Bank Account Number (may contain letters)
			//It's the account number without checksum (4 characters prefix))
			bban := account[4:accountLength]
			checksumDigits := account[2:4]

			finalAccount := lettersToValues(bban+country) + checksumDigits

			//change finalAccount to bigInt
			//and peform finalAccount % 97 on bigInt numbers
			dividend := new(big.Int)
			dividend.SetString(finalAccount, 10)

			divisor := new(big.Int)
			divisor.SetString("97", 10)

			reminder := new(big.Int)
			reminder.Mod(dividend, divisor)

			if reminder.String() == "1" {
				return true
			}
		}
	}
	return false
}

//FormatIBAN returns IBAN returns groups of four characters separated by a single space,
//the last group may have a variable length
//you can optionally verify if iban is valid
func FormatIBAN(iban string, verify bool) (formatted string, err error) {

	if verify {
		if !IsIBAN(iban) {
			return iban, errors.New("This is an invalid IBAN")
		}
	}

	ibanFields := strings.Fields(iban)
	iban = strings.Join(ibanFields, "")

	for i := 0; i < len(iban); i++ {
		if i >= 4 && i%4 == 0 {
			formatted += " "
		}
		formatted += string(iban[i])
	}
	return formatted, nil

}
