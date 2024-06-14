package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0005BruteForce(t *testing.T) {

	expectedSolution := 232792560
	bottom := 1
	top := 20
	startFrom := expectedSolution - 1000

	testname := fmt.Sprintf("Problem0005BruteForce(%d, %d) => %v \n", bottom, top, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0005BruteForce(bottom, top, startFrom)
		assert.Equal(t, expectedSolution, ans)
	})
}

func TestProblem0005Basic(t *testing.T) {

	expectedSolution := 2520
	bottom := 1
	top := 10

	testname := fmt.Sprintf("Problem0005(%d, %d) => %v \n", bottom, top, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0005(bottom, top)
		assert.Equal(t, expectedSolution, ans)
	})
}

func TestProblem0005Full(t *testing.T) {

	expectedSolution := 232792560
	bottom := 1
	top := 20

	testname := fmt.Sprintf("Problem0005(%d, %d) => %v \n", bottom, top, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0005(bottom, top)
		assert.Equal(t, expectedSolution, ans)
	})
}
