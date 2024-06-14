package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0007(t *testing.T) {

	expectedSolution := 13
	top := 6

	testname := fmt.Sprintf("Problem0007(%d) => %v \n", top, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		answer := Problem0007(top)
		assert.Equal(t, expectedSolution, answer)
	})
}
