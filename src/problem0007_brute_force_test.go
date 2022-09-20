//go:build bruteforce
// +build bruteforce

/**
 * 10001st prime
 *
 * https://projecteuler.net/problem=7
 *
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
 * we can see that the 6th prime is 13.
 *
 * What is the 10 001st prime number?
 *
 */

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
