package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsTableDriven(t *testing.T) {

	var tests = []struct {
		input int
		want  int
	}{
		{input: 1, want: 1},
		{input: 0, want: 0},
		{input: -1, want: 1},
		{input: 98765, want: 98765},
		{input: -12345, want: 12345},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Abs(%d) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := Abs(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}
