package hackerrank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifference(t *testing.T) {

	expectedSolution := 15
	input := [][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	testname := fmt.Sprintf("solveMeFirst(%d) => %v \n", input, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := DiagonalDifference(input)
		assert.Equal(t, expectedSolution, ans)
	})
}
