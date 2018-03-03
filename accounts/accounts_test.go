package accounts

import (
	"sync"
	"testing"
)

/*	-count n
Run each test and benchmark n times (default 1).
If -cpu is set, run n times for each GOMAXPROCS value.
*/

//go test -v
//go test -v -count 20 -benchmem -bench=.
//go test -v -count 20 -cpu 1,8 -benchmem -bench=.
func TestIBAN(t *testing.T) {

	var tests = []struct {
		input    string
		expected bool
	}{
		{"DE89 3704 0044 0532 0130 00 ", true},
		{"PL60102010260000042270201111", true},
		{"GB29 NWBK 6016 1331 9268 19", true},
		{"GB29 RBOS 6016 1331 9268 19", false},
		{"AT61 1904 3002 3457 3201", true},
		{"QA58 DOHB 0000 1234 5678 90AB CDEF G", true},
		{"MT84 MALT 0110 0001 2345 MTLC AST0 01S", true},
		{"", false},
		{"definitely not even close to IBAN ", false},
	}

	for _, test := range tests {
		if got := IsIBAN(test.input); got != test.expected {
			t.Errorf("IsIBAN(%q) = %v", test.input, got)
		}
	}
}

//go test -bench=.
func BenchmarkIBAN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsIBAN("MT84 MALT 0110 0001 2345 MTLC AST0 01S")
	}
}

//cpu value depends on your processor architecture
//go test -count 3 -cpu 1,2,4,6,8 -benchmem -bench=.
//The default value of cpu is the current value of GOMAXPROCS
var wg sync.WaitGroup

func BenchmarkConcIBAN(b *testing.B) {
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			IsIBAN("MT84 MALT 0110 0001 2345 MTLC AST0 01S")
		}()
	}
	wg.Wait()
}
