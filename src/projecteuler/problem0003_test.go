/**
 * Largest prime factor
 *
 * https://projecteuler.net/problem=3
 *
 * The prime factors of 13195 are 5, 7, 13 and 29.
 *
 * What is the largest prime factor of the number 600851475143 ?
 */

package projecteuler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem0003(t *testing.T) {

	solutionFound := 6857
	top := 600851475143

	testname := fmt.Sprintf("Problem0003(%d) => %v", top, solutionFound)
	t.Run(testname, func(t *testing.T) {
		ans := Problem0003(top)
		assert.Equal(t, solutionFound, ans)
	})
}
