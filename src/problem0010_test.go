/**
 * Summation of primes
 *
 * https://projecteuler.net/problem=10
 *
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 *
 * Find the sum of all the primes below two million.
 */

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0010(t *testing.T) {

	const expectedSolution = 17
	const bottom = 1
	const top = 10

	testname := fmt.Sprintf("Problem0010(%d, %d) => %v \n", bottom, top, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0010(bottom, top)
		assert.Equal(t, expectedSolution, ans)
	})
}
