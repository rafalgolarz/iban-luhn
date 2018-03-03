package cards

import (
	"sync"
	"testing"
)

//go test -v
func TestLuhn(t *testing.T) {

	var tests = []struct {
		input    string
		expected bool
	}{
		{"378282246310005", true},
		{"378282246310000", false},
		{"6011000990139424", true},
		{"", false},
		{"definitely not even close to a card number ", false},
	}

	for _, test := range tests {
		if got := IsLuhn(test.input); got != test.expected {
			t.Errorf("IsLuhn(%q) = %v", test.input, got)
		}
	}
}

//go test -bench=.
func BenchmarkLuhn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsLuhn("378282246310005")
	}
}

//cpu value depends on your processor architecture
//go test -count 3 -cpu 1,2,4,6,8 -benchmem -bench=.
//The default value of cpu is the current value of GOMAXPROCS
var wg sync.WaitGroup

func BenchmarkConcLuhn(b *testing.B) {
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			IsLuhn("378282246310005")
		}()
	}
	wg.Wait()
}
