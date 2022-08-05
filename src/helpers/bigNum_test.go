package helpers

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBigSumManyBasic(t *testing.T) {

	input := []string{"1", "2", "3"}
	expected := "6"

	ans := BigSumMany(input)

	if ans != expected {

		t.Errorf("BigSumMany(%v) = %v; want %s", input, ans, expected)
	}
}

func TestBigSumManyError(t *testing.T) {

	input := []string{"1", "2", "a"}
	expected := "0"

	ans := BigSumMany(input)

	if ans != expected {

		t.Errorf("BigSumMany(%v) = %v; want %s", input, ans, expected)
	}
}

func TestBigFactorialTableDriven(t *testing.T) {

	var tests = []struct {
		input int
		want  *big.Int
	}{
		{input: 1, want: big.NewInt(int64(1))},
		{input: 4, want: big.NewInt(int64(24))},
		{input: 5, want: big.NewInt(int64(120))},
		{input: 10, want: big.NewInt(int64(3628800))},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Abs(%d) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := BigFactorial(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}
