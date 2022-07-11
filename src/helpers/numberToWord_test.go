package helpers

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberToWordTableDriven(t *testing.T) {
	var tests = []struct {
		input big.Int
		want  string
	}{
		{input: *big.NewInt(1), want: "one"},
		{input: *big.NewInt(16), want: "sixteen"},
		{input: *big.NewInt(30), want: "thirty"},
		{input: *big.NewInt(64), want: "sixty-four"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("NumberToWord(%v) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := NumberToWord(tt.input)
			assert.Equal(t, tt.want, ans)
		})

	}
}

func BenchmarkNumberToWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberToWord(*big.NewInt((111)))
	}
}
