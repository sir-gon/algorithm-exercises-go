package projecteuler

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	in := 2
	want := true
	ans := IsPrime(2)

	if ans != want {

		t.Errorf("IsPrime(%d) = %t; want %t", in, ans, want)
	}
}

func TestIsPrimeTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{7, true},
		{13, true},
		{29, true},

		{4, false},
		{10, false},
		{100, false},
		{3000, false},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("IsPrime(%d) => %t", tt.a, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := IsPrime(tt.a)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(1)
	}
}
