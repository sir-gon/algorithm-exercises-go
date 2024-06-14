package hackerrank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiniMaxSumBorderCase(t *testing.T) {
	expectedSolution := []int{1, 1}
	input := []int{}

	testname := fmt.Sprintf("MiniMaxSum(%v) => %d \n", input, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		_, err := MiniMaxSum(input)
		assert.Error(t, err)
	})
}

func TestMiniMaxSum(t *testing.T) {
	var tests = []struct {
		input []int
		want  string
	}{
		{input: []int{1, 2, 3, 4, 5}, want: "10 14"},
		{input: []int{5, 4, 3, 2, 1}, want: "10 14"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("MiniMaxSum(%v) => %s", tt.input, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans, _ := MiniMaxSum(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}
