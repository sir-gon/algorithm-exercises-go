//go:build bruteforce
// +build bruteforce

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

	expectedSolution := 142913828922

	testname := fmt.Sprintf("Problem0010() => %v \n", expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0010()
		assert.Equal(t, expectedSolution, ans)
	})
}
