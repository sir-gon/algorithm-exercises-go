/**
 *
 * https://projecteuler.net/problem=XX
 *
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
 * we can see that the 6th prime is 13.
 *
 * What is the 10 001st prime number?
 *
 * ////////////////////////////////////////////////////////////////////////////
 * Solution found: 104743
 * ////////////////////////////////////////////////////////////////////////////
 */

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0007(t *testing.T) {

	expectedSolution := 104743
	bottom := 1
	top := 10001

	testname := fmt.Sprintf("Problem0000Template(%d, %d) => %v \n", bottom, top, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		answer := Problem0007(bottom, top)
		assert.Equal(t, expectedSolution, answer)
	})
}
