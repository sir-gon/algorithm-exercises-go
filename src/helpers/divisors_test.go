package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisorsBasic(t *testing.T) {
	ans := Divisors(1)

	if len(ans) != 1 {

		t.Errorf("Divisors() = %d; want 1", ans)
	}
}

func TestDivisorsBasicTableDriven(t *testing.T) {

	var tests = []struct {
		input int
		want  []int
	}{
		{input: 1, want: []int{1}},
		{input: 2, want: []int{1, 2}},
		{input: 10, want: []int{1, 2, 5, 10}},
		{input: 16, want: []int{1, 2, 4, 8, 16}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Divisors(%d) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := Divisors(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}
