package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeTableDriven(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{7, true},
		{13, false},
		{29, false},

		{123, false},
		{534, false},
		{101, true},
		{9889, true},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("IsPalindrome(%d) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := IsPalindrome(tt.input)
			assert.Equal(t, tt.want, ans)
		})

	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome(111)
	}
}
