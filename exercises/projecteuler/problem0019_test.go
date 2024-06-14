package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gon.cl/algorithms/exercises/projecteuler/constants"
)

func TestProblem0019Small(t *testing.T) {

	const expectedSolution = 2

	const sinceInput = 1900
	const untilInput = 1901

	testname := fmt.Sprintf("Problem0019(%d, %d) => %v \n", sinceInput, untilInput, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0019(constants.SUNDAY, sinceInput, untilInput)
		assert.Equal(t, expectedSolution, ans)
	})
}

func TestProblem0019Full(t *testing.T) {

	const expectedSolution = 171

	const sinceInput = 1901
	const untilInput = 2000

	testname := fmt.Sprintf("Problem0019(%d, %d) => %v \n", sinceInput, untilInput, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0019(constants.SUNDAY, sinceInput, untilInput)
		assert.Equal(t, expectedSolution, ans)
	})
}
