//go:build bruteforce
// +build bruteforce

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
