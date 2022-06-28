//go:build bruteforce
// +build bruteforce

/**
 * Longest Collatz sequence
 *
 * https://projecteuler.net/problem=14
 *
 *
 * The following iterative sequence is defined for the set of positive integers:
 *
 * n → n/2 (n is even)
 * n → 3n + 1 (n is odd)
 *
 * Using the rule above and starting with 13, we generate the following sequence:
 *
 * 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
 * It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
 * Although it has not been proved yet (Collatz Problem), it is thought that all
 * starting numbers finish at 1.
 *
 * Which starting number, under one million, produces the longest chain?
 *
 * NOTE: Once the chain starts the terms are allowed to go above one million.
 */

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0014BruteForce(t *testing.T) {

	const expectedSolution = 837799
	const inputBottom = 1
	const inputTop = 1000000

	testname := fmt.Sprintf("Problem0014(%d, %d) => %d \n", inputBottom, inputTop, expectedSolution)
	t.Run(testname, func(t *testing.T) {

		ans := Problem0014(inputBottom, inputTop)
		assert.Equal(t, expectedSolution, ans)
	})
}
