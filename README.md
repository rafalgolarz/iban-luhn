# bankingo
Go libraries for validating banking data (account or card numbers)

## [accounts.go](https://github.com/rafalgolarz/bankingo/blob/master/accounts/accounts.go)

#### func IsIBAN(account string) bool

Returns true if account is a valid IBAN number
#### func FormatIBAN(iban string, verify bool) (formatted string, err error)

Returns formatted IBAN.
If verify is set to true, perform IBAN validation first (by calling IsIBAN)

## [cards.go](https://github.com/rafalgolarz/bankingo/blob/master/cards/cards.go)

## Running the example

Follow [instructions](https://golang.org/doc/install) to install go.
Once you have Go up and running, you can download, build and run the example using the following commands.

    $ go get github.com/rafalgolarz/bankingo/accounts

or

    $ go get github.com/rafalgolarz/bankingo/cards

then, assuming you have $GOPATH environmental variable set

    $ cd $GOPATH/src/github.com/rafalgolarz/bankingo
    $ go run example.go
    
## Running tests

    $ cd $GOPATH/src/github.com/rafalgolarz/bankingo/accounts
    $ go test -v
    $ go test -bench=.
    
    $ cd $GOPATH/src/github.com/rafalgolarz/bankingo/cards
    $ go test -v
    $ go test -bench=.
    