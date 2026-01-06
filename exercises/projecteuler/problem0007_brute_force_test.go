//go:build bruteforce

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0007BruteForce(t *testing.T) {

	expectedSolution := 104743
	top := 10001

	testname := fmt.Sprintf("Problem0007(%d) => %v \n", top, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		answer := Problem0007(top)
		assert.Equal(t, expectedSolution, answer)
	})
}
