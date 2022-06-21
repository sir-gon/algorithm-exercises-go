package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductDriven(t *testing.T) {

	var tests = []struct {
		input []int
		want  int
	}{
		{input: []int{1, 2, 4, 8, 0}, want: 0},
		{input: []int{1, 2, 3, 4}, want: 24},
		{input: []int{-1, -2, 1, 2}, want: 4},
		{input: []int{}, want: 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("Product(%d) => %v", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := Product(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}
