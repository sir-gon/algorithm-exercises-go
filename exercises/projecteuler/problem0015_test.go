package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0015(t *testing.T) {

	expectedSolution := 137846528820
	inputGridSize := 20

	testname := fmt.Sprintf("Problem0015(%d) => %v \n", inputGridSize, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0015(inputGridSize)
		assert.Equal(t, expectedSolution, ans)
	})
}
