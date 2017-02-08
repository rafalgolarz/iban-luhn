package main

import (
	"fmt"

	"github.com/rafalgolarz/bankingo/accounts"
	"github.com/rafalgolarz/bankingo/cards"
)

func main() {

	fmt.Println(cards.IsLuhn("3782 8224 6310 005"))
	fmt.Println(accounts.IsIBAN("GB29 NWBK 6016 1331 9268 19"))

	//format the input string as IBAN
	formatted, _ := accounts.FormatIBAN("MT84MALT011000012345MTLCAST001S", false)
	fmt.Println(formatted)

	//set IBAN validation flag to true
	result, err := accounts.FormatIBAN("MT84MALT011000012345MTLCAST002S", true)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
