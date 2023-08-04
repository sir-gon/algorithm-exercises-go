package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0012(t *testing.T) {

	expectedSolution := 28
	inputTop := 5

	testname := fmt.Sprintf("Problem0012(%d) => %v \n", inputTop, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0012(inputTop)
		assert.Equal(t, expectedSolution, ans)
	})
}
