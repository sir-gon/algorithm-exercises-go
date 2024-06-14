package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0017(t *testing.T) {

	const expectedSolution = 21124
	const inputInit = 1
	const inputLast = 1000

	testname := fmt.Sprintf("Problem0017(%d, %d) => %v \n", inputInit, inputLast, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0017(inputInit, inputLast)
		assert.Equal(t, expectedSolution, ans)
	})
}

func TestProblem0017ErrorBounds(t *testing.T) {

	const expectedSolution = -1
	const inputInit = 999
	const inputLast = 1001

	testname := fmt.Sprintf("Problem0017(%d, %d) => %v \n", inputInit, inputLast, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0017(inputInit, inputLast)
		assert.Equal(t, expectedSolution, ans)
	})
}
