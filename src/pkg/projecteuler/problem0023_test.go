package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0023SmallCase(t *testing.T) {

	const expectedSolution = 276
	const inputUnderLimit = 12
	const inputSuperLimit = 24

	testname := fmt.Sprintf("Problem0023(%d, %d) => %v \n", inputUnderLimit, inputSuperLimit, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0023(inputUnderLimit, inputSuperLimit)
		assert.Equal(t, expectedSolution, ans)
	})
}

func TestProblem0023(t *testing.T) {

	const expectedSolution = 4179871
	const inputUnderLimit = 12
	const inputSuperLimit = 28123

	testname := fmt.Sprintf("Problem0023(%d, %d) => %v \n", inputUnderLimit, inputSuperLimit, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0023(inputUnderLimit, inputSuperLimit)
		assert.Equal(t, expectedSolution, ans)
	})
}
