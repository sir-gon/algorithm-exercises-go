/**
 * Multiples of 3 and 5
 *
 * https://projecteuler.net/problem=1
 *
 * If we list all the natural numbers below 10 that are multiples
 * of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 *
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */

package projecteuler

import (
	"testing"
)

func TestProblem0001(t *testing.T) {

	solutionFound := 233168
	top := 1000

	ans := Problem0001(top)

	if ans != solutionFound {
		t.Errorf("Problem0001(%d) = %d; want %d", top, ans, solutionFound)
	}
}
