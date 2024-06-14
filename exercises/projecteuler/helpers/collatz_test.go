package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollatzTableDriven(t *testing.T) {

	var tests = []struct {
		input int
		want  int
	}{
		{input: 13, want: 40},
		{input: 40, want: 20},
		{input: 20, want: 10},
		{input: 10, want: 5},
		{input: 5, want: 16},
		{input: 16, want: 8},
		{input: 8, want: 4},
		{input: 4, want: 2},
		{input: 2, want: 1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Collatz(%d) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := Collatz(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}

func TestCollatzSequenceTableDriven(t *testing.T) {

	var tests = []struct {
		input int
		want  []int
	}{
		{input: 13, want: []int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("CollatzSequence(%d) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := CollatzSequence(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}
