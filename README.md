[![Build Status](https://travis-ci.org/rafalgolarz/iban-luhn.svg?branch=master)](https://travis-ci.org/rafalgolarz/iban-luhn)
[![Go Report Card](https://goreportcard.com/badge/github.com/rafalgolarz/iban-luhn)](https://goreportcard.com/report/github.com/rafalgolarz/iban-luhn)
[![codebeat badge](https://codebeat.co/badges/3cadc60b-3642-46bc-9118-1595e354aa6d)](https://codebeat.co/projects/github-com-rafalgolarz-iban-luhn)
[![CircleCI](https://circleci.com/gh/rafalgolarz/iban-luhn/tree/master.svg?style=svg)](https://circleci.com/gh/rafalgolarz/iban-luhn/tree/master)
[![GoDoc](https://godoc.org/github.com/rafalgolarz/iban-luhn?status.svg)](https://godoc.org/github.com/rafalgolarz/iban-luhn)

# iban-luhn
Go libraries for validating banking data (account or card numbers)

## [accounts.go](https://github.com/rafalgolarz/iban-luhn/blob/master/accounts/accounts.go)

#### func IsIBAN(account string) bool

Returns true if account is a valid IBAN number

#### func FormatIBAN(iban string, verify bool) (formatted string, err error)

Returns formatted IBAN.
If verify is set to true, perform IBAN validation first (by calling IsIBAN)

## [cards.go](https://github.com/rafalgolarz/iban-luhn/blob/master/cards/cards.go)

#### func IsLuhn(number string) bool

Returns true if number passes the Luhn validation

## Running the example

Follow [instructions](https://golang.org/doc/install) to install go.
Once you have Go up and running, you can download, build and run the example using the following commands.

    $ go get github.com/rafalgolarz/iban-luhn/accounts

or

    $ go get github.com/rafalgolarz/iban-luhn/cards

then, assuming you have $GOPATH environmental variable set

    $ cd $GOPATH/src/github.com/rafalgolarz/iban-luhn
    $ go run example.go
    
## Running tests

    $ cd $GOPATH/src/github.com/rafalgolarz/iban-luhn/accounts
    $ go test -v
    $ go test -bench=.
    $ go test -count 4 -benchmem -bench=.
    
    $ cd $GOPATH/src/github.com/rafalgolarz/bankingo/cards
    $ go test -v
    $ go test -bench=.
    $ go test -count 4 -benchmem -bench=.
