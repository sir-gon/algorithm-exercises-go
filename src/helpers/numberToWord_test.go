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
		{input: *big.NewInt(20), want: "twenty"},
		{input: *big.NewInt(30), want: "thirty"},
		{input: *big.NewInt(64), want: "sixty-four"},
		{input: *big.NewInt(301), want: "three hundred and one"},
		{input: *big.NewInt(348), want: "three hundred and forty-eight"},
		{input: *big.NewInt(500), want: "five hundred"},
		{input: *big.NewInt(1000), want: "one thousand"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("NumberToWord(%v) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans, errOutOfBounds := NumberToWord(tt.input)
			if !errOutOfBounds {
				assert.Equal(t, tt.want, ans)
			}
		})

	}
}

func TestNumberToWordBoundsTableDriven(t *testing.T) {
	var tests = []struct {
		input big.Int
		want  string
	}{
		{input: *big.NewInt(0), want: "zero"},
		{input: *big.NewInt(20), want: "twenty"},
		{input: *big.NewInt(99), want: "ninety-nine"},
		{input: *big.NewInt(100), want: "one hundred"},
		{input: *big.NewInt(999), want: "nine hundred and ninety-nine"},
		{input: *big.NewInt(1000), want: "one thousand"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("NumberToWord(%v) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans, errOutOfBounds := NumberToWord(tt.input)
			if !errOutOfBounds {
				assert.Equal(t, tt.want, ans)
			}
		})

	}
}

func TestNumberToWordErrorTableDriven(t *testing.T) {
	var tests = []struct {
		input big.Int
		want  string
	}{
		{input: *big.NewInt(1001), want: ""},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("NumberToWord(%v) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans, errOutOfBounds := NumberToWord(tt.input)
			if errOutOfBounds {
				assert.Equal(t, tt.want, ans)
			}
		})

	}
}

func BenchmarkNumberToWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberToWord(*big.NewInt((111)))
	}
}
