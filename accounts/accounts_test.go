package accounts

import "testing"

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

func BenchmarkIBAN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsIBAN("MT84 MALT 0110 0001 2345 MTLC AST0 01S")
	}
}
